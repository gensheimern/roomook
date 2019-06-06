<template>
  <v-content>
    <v-container fluid fill-height>
      <transition name="bounce">
        <v-layout>
          <v-fab-transition>
            <v-btn
              id="CreateButtonFloating"
              v-show="!hidden"
              color="#274d6f"
              fab
              dark
              absolute
              top
              left
              @click="toggleCreateBookingModal"
            >
              <v-icon>add</v-icon>
            </v-btn>
          </v-fab-transition>
          <v-flex>
            <v-sheet height="100%">
              <v-calendar
                ref="calendar"
                v-model="start"
                :now="today"
                :type="type"
                :first-interval="firstInterval"
                :interval-count="intervalCount"
                :interval-height="intervalHeight"
                :weekdays="weekdays"
                @click:interval="test($event)"
              >
                <template v-slot:dayHeadere="{ date }"></template>

                <template v-slot:dayBody="{ date, timeToY, minutesToPixels } ">
                  <template v-for="event in eventsMap[date]">
                    <div
                      v-if="event.time"
                      :key="event.eventID"
                      :style="{ top: timeToY(event.time) + 'px', height: minutesToPixels(event.duration) + 'px' , backgroundColor: randomBackgroundColor()}"
                      class="my-event with-time"
                      @click="toggleBookingDetails(event)"
                      v-html="event.title"
                    ></div>
                  </template>
                </template>
              </v-calendar>
            </v-sheet>
          </v-flex>
          <CreateWrapper
            ref="createWrapper"
            @loadingWhileCreateWrapper="loadingWhileCreateWrapper"
          />

          <BookingDetailCard
            v-if="bookingDetailCard"
            :style="{top: topStyle, left: leftStyle, width: widthStyle}"
            class="bookingDetails"
            :room="room"
            :clickedEvent="clickedEventDetail"
            :activeRoomID="activeRoomID"
            :loggedInUser="loggedInUser"
            @closeDetailCardHook="closeDetailCardHook"
          />
        </v-layout>
      </transition>
    </v-container>
  </v-content>
</template>

<script>
import CreateWrapper from "./CreateWrapper.vue";
import BookingDetailCard from "./BookingDetailCard.vue";

var today = new Date();

export default {
  props: {
    AllBookings: { type: Array },
    calendarNextPage: { type: Function },
    activeRoomID: { type: String },
    loggedInUser: { type: String }
  },
  data: () => ({
    topStyle: "50%",
    widthStyle: "400px",
    leftStyle: "70%",

    type: "week",
   

    firstInterval: "7",
    intervalCount: "13",
    intervalHeight: "60",
    weekdays: [1, 2, 3, 4, 5],

    today: today.toISOString().substring(0, 10),
    start: today.toISOString().substring(0, 10),
    loading: false,
    bookingDetailCard: false,
    hidden: false,
    events: [],
    reload: false,
    clickedEventDetail: {},
    openCreate: false,
    backgroundColors: ["	 #94b8b8"],
    room: null
  }),
  components: {
    BookingDetailCard,
    CreateWrapper
  },

  computed: {
    // convert the list of events into a map of lists keyed by date
    eventsMap() {
      this.events = [];
      if (this.AllBookings != null && this.AllBookings != "") {
        this.AllBookings.forEach(item => {
          var beginDate = new Date(item.begin[0]);
          var endDate = new Date(item.end[0]);
          var bookingDuration = this.timeDiff(beginDate, endDate);
          this.events.push({
            eventID: item.bookingID,
            title: item.title,
            details: "Going to e beach!",
            duration: bookingDuration,
            date: item.begin[0].substr(0, 10),
            time: item.begin[0].substr(10, 15),
            end: item.end[0].substr(10, 15),
            open: false,
            recurringType: item.recurringType,
            parentID: item.parentID,
            roomID: item.roomID,
            owner: item.userID,
            baseURL: process.env
          });
        });
      }
      const map = {};
      this.events.forEach(e => (map[e.date] = map[e.date] || []).push(e));
      return map;
    }
  },

  created() {
    if (window.innerWidth <= 500) {
      this.type = "day";
      this.leftStyle = "10%";
      this.topStyle = "30%";
      this.widthStyle = "300px";
    }

    
  },

  mounted(){
console.log(this.baseURL)
  },

  methods: {
    randomBackgroundColor() {
      var color = this.backgroundColors[
        Math.floor(Math.random() * Math.floor(this.backgroundColors.length))
      ];
      return color;
    },
    loadingWhileCreateWrapper(booking) {
      this.$emit("fetchBookingByRoom", booking);
    },
    closeDetailCardHook() {
      this.bookingDetailCard = 0;
    },
    toggleBookingDetails(event) {
      if (
        this.bookingDetailCard != 0 &&
        this.bookingDetailCard == event.eventID
      ) {
        this.bookingDetailCard = 0;
      } else if (
        this.bookingDetailCard != 0 &&
        this.bookingDetailCard != event.eventID
      ) {
        this.clickedEventDetail = event;
        this.bookingDetailCard = event.eventID;
      } else if (this.bookingDetailCard == 0) {
        this.clickedEventDetail = event;
        this.bookingDetailCard = event.eventID;
      }
    },

    timeDiff(beginTime, endTime) {
      var one_minute = 1000 * 60;

      var begin_min = beginTime.getTime();
      var end_min = endTime.getTime();

      var diff = end_min - begin_min;

      return Math.round(diff / one_minute);
    },
    nextPage() {
      this.$refs.calendar.next();
    },
    previousPage() {
      this.$refs.calendar.prev();
    },
    toggleCreateOpen() {
      if (this.openCreate == false) {
        this.openCreate = true;
      } else {
        this.openCreate = false;
      }
    },
    toggleCreateBookingModal() {
      this.$refs.createWrapper.toggleModal();
    },

    setRoom(room) {
      console.log(room, "raum");
      this.room = room;
      this.$refs.createWrapper.setRoom(room);
    },

    changeCalendarStartDate(startDate) {
      this.start = startDate;
    },
    changeCalendarType(type) {
      this.type = type;
    }
  }
};
</script>

<style  scoped>
.my-event {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  border-radius: 2px;
  background-color: #9e9690;
  color: #ffffff;
  border: 1px solid #9e9690;
  font-size: 12px;
  padding: 3px;
  cursor: pointer;
  margin-bottom: 1px;
  left: 4px;
  margin-right: 8px;
  position: relative;
}
.with-time {
  position: absolute;
  right: 4px;
  margin-right: 0px;
}

#CreateButtonFloating {
  margin-top: 5vh;
}

.bookingDetails {
  height: auto;
  position: fixed;
}
</style>