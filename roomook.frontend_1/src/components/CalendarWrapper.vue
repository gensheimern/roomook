<template v-if="isLoading && x">
  <v-app id="inspire">
    <CalendarDrawer
      ref="SideDrawer"
      :loggedInUser="loggedInUser"
      @fetchBookingByRoom="fetchBookingByRoom"
      @changeCalendarStartDate="changeCalendarStartDate"
      @changeCalendarType="changeCalendarType"
      @toggleRoomView="toggleRoomView"
      @refreshFreeRooms="refreshFreeRooms"
    />
    <v-content v-if="room==null && !roomView ">
      <v-flex class="noRoomSelected" xs12>
        <v-flex>
          <h1 style="font-size:40px;color:#C0C0C0">Please Select a Room</h1>
        </v-flex>
        <v-flex>
          <span
            style="color:#C0C0C0"
          >You also can set your favorite room which will be displayed as default</span>
        </v-flex>
      </v-flex>
    </v-content>
    <HeaderCalendar
      v-if="!roomView"
      class="Header"
      ref="headerCalendar"
      @toggleDrawer="toggleDrawer"
      @changeCalendarType="changeCalendarType"
      @changeCalendarStartDate="changeCalendarStartDate"
      @calendarPreviousPage="calendarPreviousPage"
      @calendarNextPage="calendarNextPage"
      @fetchBookingByRoom="fetchBookingByRoom"
      @toggleRoomView="toggleRoomView"
      @refreshFreeRooms="refreshFreeRooms"
    />
    <transition name="slide-fade">
      <Calendar
        v-if="!roomView &&room!=null"
        class="Calendar"
        @fetchBookingByRoom="fetchBookingByRoom"
        ref="calendar"
        :activeRoomID="activeRoom"
        :AllBookings="bookings"
        :loggedInUser="loggedInUser"
      />
    </transition>

    <HeaderFindRoom v-if="roomView" @toggleDrawer="toggleDrawer"/>
    <transition name="slide-fade">
      <FreeRoomWrapper
        v-if="roomView"
        :freeRooms="freeRooms"
        @getFreeRooms="getFreeRooms"
        @fetchBookingByRoom="fetchBookingByRoom"
        @refreshFreeRooms="refreshFreeRooms"
        @toggleRoomView="toggleRoomView"
      />
    </transition>
  </v-app>
</template>



<script>
import HeaderCalendar from "./HeaderCalendar.vue";
import HeaderFindRoom from "./HeaderFindRoom.vue";
import Calendar from "./Calendar.vue";
import FreeRoomWrapper from "./FreeRoomWrapper.vue";
import CalendarDrawer from "./CalendarDrawer.vue";
import axios from "axios";

export default {
  name: "CalendarWrapper",
  data() {
    return {
      isLoading: false,
      roomView: false,
      bookings: [],
      nextWeek: "",
      room: null,
      favRoom: this.$cookies.get("favoriteRoom"),
      activeRoom: null,
      ws: null,
      loggedInUser: null,
      freeRooms: [],
      isLoggedIn: false,
      x: false
    };
  },
  components: {
    HeaderCalendar,
    HeaderFindRoom,
    Calendar,
    FreeRoomWrapper,
    CalendarDrawer
  },
  watch: {
    isLoggedIn: function() {
      this.setUpComponent();
    }
  },
  created() {
    this.isLoggedIn = localStorage.getItem("Authentication");
    if (process.env.VUE_APP_AUTH_AS_MIDDLEWARE != "0") {
      axios.defaults.headers.common["Authorization"] =
        "Bearer " + localStorage.getItem("Authentication");
    }
    axios.defaults.headers.common["Content-Type"] = "application/json";
  },

  mounted() {
    this.$refs.headerCalendar.changeRoomName(this.favRoom.name);
  },
  beforeRouteUpdate(to, from, next) {
    this.setUpComponent();
    next();
  },

  methods: {
    setUpComponent() {
      this.$options.sockets.onmessage = data => {
        console.log(data ,"data")
        if (
          JSON.parse(data.data).Type == "createBooking" &&
          JSON.parse(data.data).RoomID == this.activeRoom
        ) {
          axios
            .get(this.$url + "/b/r/" + JSON.parse(data.data).RoomID)
            .then(response => {
              if (response.status === 200) {
                this.bookings = response.data;
              }
            })
            .catch(error => console.log(error));
        }
      };
      axios
        .get(this.$url + "/user/loggedinuser/")
        .then(response => {
          if (response.status === 200) {
            this.loggedInUser = response.data.user;
          }
        })
        .catch(err => {
          if (err.response.status === 401) {
            this.$router.push("/login");
          }
          console.log(err);
        });
      this.fetchBooking(this.favRoom);
    },
    getFreeRooms(beginDate, endDate) {
      this.$refs.SideDrawer.getFreeRooms(beginDate, endDate);
    },
    fetchBooking(room) {
      this.room = room;
      this.bookings = [];
      this.activeRoom = room.id;

      axios
        .get(this.$url + "/b/r/" + room.id)
        .then(async response => {
          console.log(response.status);
          if (response.status === 200) {
            this.bookings = response.data;
            await this.$nextTick();
            this.$refs.calendar.setRoom(room);
          }
        })
        .catch(error => {
          console.log(error);

          if (error.response.status === 401) {
            this.$router.push("/login");
          }
        });
    },
    createBookingWithWebsocket() {},
    calendarNextPage() {
      this.$refs.calendar.nextPage();
    },
    calendarPreviousPage() {
      this.$refs.calendar.previousPage();
    },
    async fetchBookingByRoom(room) {
      this.roomView = false;

      this.fetchBooking(room);

      //Sets the Room Name in the Header to the selected room
      //Pause until the Components are rendered in template so we can use the reference
      await this.$nextTick();
      this.$refs.headerCalendar.changeRoomName(room.name);
      //Passes the selected room to Calendar component
      this.$refs.calendar.setRoom(room);
    },
    changeCalendarStartDate(startDate) {
      this.$refs.calendar.changeCalendarStartDate(startDate);
    },
    changeCalendarType(type) {
      this.$refs.calendar.changeCalendarType(type);
    },
    toggleDrawer() {
      this.$refs.SideDrawer.toggleDrawer();
    },

    refreshFreeRooms(freeRooms) {
      this.freeRooms = freeRooms;
    },
    toggleRoomView(toggle) {
      if (toggle == true) {
        this.roomView = true;
      } else if (toggle == false) {
        this.roomView = false;
      }
    }
  }
};
</script>

<style scope>
.v-calendar-daily__scroll-area {
  overflow: hidden !important;
}
/* Enter and leave animations can use different */
/* durations and timing functions.              */
.slide-fade-enter-active {
  /* This timing applies on the way OUT */
  transition-timing-function: ease-in;

  /* Quick on the way out */
  transition: 0.2s;

  /* Hide thing by pushing it outside by default */
  transform: translateX(130%);
}

.noRoomSelected {
  margin-top: 40vh;
}
</style>

