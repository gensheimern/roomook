<template>
  <BookingTemplate
    ref="BookingTemplate"
    :timeBegin="timeBegin"
    :timeEnd="timeEnd"
    :date="date"
    :selectRoom="selectRoom"
    :whoIsCalling="whoIsCalling"
    :room="room"
    :title="title"
    :showRecurrently="showRecurrently"
    @AjaxRequest="createBookingAjax"
    @closeCreateModal="closeCreateModal"
  />
</template>

<script>
import axios from "axios";
import { setTimeout } from "timers";
import BookingTemplate from "./BookingTemplate";
var moment = require("moment");

const pause = ms => new Promise(resolve => setTimeout(resolve, ms));

export default {
  data: vm => ({
    valid: false,
    whoIsCalling: "CreateBooking",
    title: "",

    selectRoom: null,
    showRecurrently: true,
    recurringType: "",

    date: new Date().toISOString().substr(0, 10),
    timeBegin: moment().format("HH") + ":" + moment().format("mm"),
    timeEnd: "",

    room: null,
    booking: []
  }),

  computed: {
    fields() {
      if (!this.model) return [];

      return Object.keys(this.model).map(key => {
        return {
          key,
          value: this.model[key] || "n/a"
        };
      });
    }
  },

  created() {
    if (this.$cookies.get("favoriteRoom") == null) {
      this.selectRoom = this.room;
    } else {
      this.room = this.$cookies.get("favoriteRoom").name;
      this.selectRoom = this.$cookies.get("favoriteRoom").name;
    }
  },

  methods: {
    UpdateTime() {
      this.timeEnd =
        moment()
          .add(30, "minutes")
          .format("HH") +
        ":" +
        moment()
          .add(30, "minutes")
          .format("mm");

      this.timeBegin = moment().format("HH") + ":" + moment().format("mm");
    },

    closeCreateModal() {
      this.$emit("toggleModal");
    },

    setRoom(room) {
      this.$refs.BookingTemplate.setRoom(room);
    },

    // AJAX Request to create bookings
    createBookingAjax(datetimeBegin, datetimeEnd, title, recurringType) {
      //Filling Date Arrays

      axios
        .post(this.$url + "/booking/", {
          Title: title,
          RoomID: this.$refs.BookingTemplate.getSearchObject().id,
          Begin: datetimeBegin,
          End: datetimeEnd,
          RecurringType: recurringType
        })
        .then(async response => {
          if (response.status === 401) {
            this.$router.push("/login");
          } else if (response.status === 201) {
            this.$refs.BookingTemplate.alert(true, "success", "Created");

            this.booking = response.data;
            this.$emit("loadingWhileCreate", {
              id: this.$refs.BookingTemplate.getSearchObject().id,
              name: this.$refs.BookingTemplate.getSearchObject().name
            });
            await pause(1000);

            this.$emit("toggleModal");
            this.$refs.BookingTemplate.alert(false, "success", "Created");

            this.$socket.sendObj({
              Type: "createBooking",
              RoomID: this.$refs.BookingTemplate.getSearchObject().id
            });
          }
        })
        .catch(error => {
          this.$refs.BookingTemplate.alert(
            true,
            "error",
            error.response.data.error
          );
          console.log(error);
        });
    }
  },
  components: {
    BookingTemplate
  }
};
</script>

<style scoped>
.headline {
  background-color: #274d6f;
  width: 100%;
}
.v-card__title {
  padding: 0px !important;

  color: white;
}

.createCardImage {
  width: 100%;
  height: auto;
}
</style>