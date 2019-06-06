<template>
  <div class="Header">

    <v-toolbar v-if="prevWeekfirstDay && nextWeekfirstDay && prevWeeklastDay && prevWeeklastDay"  fixed app dark color="#274d6f">
      <v-toolbar-side-icon @click.stop="toggleDrawer()" font-color="white"></v-toolbar-side-icon>

      <v-flex 4xs v-if="!isMobile">
        <v-btn flat fab outline small color="white" @click="calendarPreviousPage">
          <v-icon dark>keyboard_arrow_left</v-icon>
        </v-btn>

        <span id="prevWeekSpan">
          {{prevWeekfirstDay.getUTCFullYear() + "."+
          ('0' + (prevWeekfirstDay.getUTCMonth() +1)).slice(-2) + "." +
          ('0' + prevWeekfirstDay.getUTCDate()).slice(-2) + " - " +
          prevWeeklastDay.getUTCFullYear() + "."+
          ('0' + (prevWeeklastDay.getUTCMonth() +1)).slice(-2) + "." +
          ('0' + prevWeeklastDay.getUTCDate()).slice(-2)
          }}
        </span>
      </v-flex>

      <v-flex 4xs>
        <div @click="selectRoom()">
          <h1>{{room}}</h1>
        </div>
      </v-flex>
      <v-flex 2xs v-if="!isMobile">
        <span id="prevWeekSpan">
          {{nextWeekfirstDay.getUTCFullYear() + "."+
          ('0' + (nextWeekfirstDay.getUTCMonth() +1)).slice(-2) + "." +
          ('0' + nextWeekfirstDay.getUTCDate()).slice(-2) + " - " +
          nextWeeklastDay.getUTCFullYear() + "."+
          ('0' + (nextWeeklastDay.getUTCMonth() +1)).slice(-2) + "." +
          ('0' + nextWeeklastDay.getUTCDate()).slice(-2) }}
        </span>

        <v-btn fab outline small color="white" @click="calendarNextPage">
          <v-icon dark>keyboard_arrow_right</v-icon>
        </v-btn>
      </v-flex>
    </v-toolbar>
  </div>
</template>


<script>
import { finished } from "stream";
var moment = require("moment");
moment().format();
export default {
  name: "HeaderCalendar",

  props: {
    roomName: { typ: String }
  },
  data() {
    return {
      isMobile: false,

      room: null,
      today: new Date(),
      prevWeekfirstDay: "",
      prevWeeklastDay: "",
      nextWeekfirstDay: "",
      nextWeeklastDay: ""
    };
  },


  created() {
    if (this.$cookies.get("favoriteRoom") == null) {
      this.room = "Select Room";
    } else {
      this.room = this.$cookies.get("favoriteRoom").name;
    }
   
    if (window.innerWidth <= 500) {
      this.isMobile = true

    }
  },
  mounted() {
    this.parsePreviousWeekSpan();
    this.parseNextWeekSpan();
  },

  methods: {
    changeRoomName(roomName) {
      this.room = roomName;
    },
    calendarNextPage() {
      this.today = moment(this.today)
        .add(7, "d")
        .toDate();
      this.$emit("calendarNextPage");
      this.parsePreviousWeekSpan();
      this.parseNextWeekSpan();
    },
    calendarPreviousPage() {
      this.today = moment(this.today)
        .subtract(7, "d")
        .toDate();

      this.$emit("calendarPreviousPage");
      this.parsePreviousWeekSpan();
      this.parseNextWeekSpan();
    },
    parsePreviousWeekSpan() {
      var prevweekfirst = new Date(
        this.today.getFullYear(),
        this.today.getMonth(),
        this.today.getDate() - 6
      );
      this.prevWeekfirstDay = new Date(
        prevweekfirst.setDate(
          prevweekfirst.getDate() - prevweekfirst.getDay() + 2
        )
      );
      var prevweeklast = new Date(
        this.today.getFullYear(),
        this.today.getMonth(),
        this.today.getDate() - 6
      );

      this.prevWeeklastDay = new Date(
        prevweeklast.setDate(prevweeklast.getDate() - prevweeklast.getDay() + 6)
      );
    },
    parseNextWeekSpan() {
      var nextweekfirst = new Date(
        this.today.getFullYear(),
        this.today.getMonth(),
        this.today.getDate() + 6
      );
      this.nextWeekfirstDay = new Date(
        nextweekfirst.setDate(
          nextweekfirst.getDate() - nextweekfirst.getDay() + 2
        )
      );
      var nextweeklast = new Date(
        this.today.getFullYear(),
        this.today.getMonth(),
        this.today.getDate() + 6
      );

      this.nextWeeklastDay = new Date(
        nextweeklast.setDate(nextweeklast.getDate() - nextweeklast.getDay() + 6)
      );
    },
    toggleDrawer() {
      this.$emit('toggleDrawer');
    },
    toggleRoomView(toggle, freeRooms) {
      this.$emit("toggleRoomView", toggle, freeRooms);
    },

    fetchBookingByRoom(room) {
      this.$emit("fetchBookingByRoom", room);
    },
    changeCalendarStartDate(startDate) {
      this.$emit("changeCalendarStartDate", startDate);
    },
    changeCalendarType(type) {
      this.$emit("changeCalendarType", type);
    }
  }
};
</script>


<style scope>
h1,
#prevWeekSpan {
  color: white;
  margin-left: 10px;
  margin-right: 10px;
  display: inline-block;
}

</style>