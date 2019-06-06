<template>
  <v-dialog
    v-model="dialog"
    width="700"
    scrollable
    transition="dialog-bottom-transition"
    hide-overlay
  >
    <template v-slot:activator="{ on }"></template>

    <v-card>
      <v-img class="createCardImage" :src="bannerURL" aspect-ratio="1.9"></v-img>

      <v-layout row wrap>
        <v-flex xs12>
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
            @AjaxRequest="updateBookingAjax"
            @closeCreateModal="closeCreateModal"
          />
        </v-flex>
      </v-layout>
    </v-card>
    <UpdateBooking  ref="UpdateBooking"  @closeCreateModal="closeCreateModal"/>
  </v-dialog>
</template>

<script>
import BookingTemplate from "./BookingTemplate";
import UpdateBooking from "./UpdateBooking";
import moment from "moment";
import axios from "axios";
export default {
  props: {
    activeRoomID: { type: String },
    booking: { type: Object },
    room: { type: Object }
  },
  data() {
    return {
      dialog: false,
      showRecurrently: false,

      timeBegin: "",
      timeEnd: "",
      timeEnd: "",
      date: null,
      title: "",
      selectRoom: null,
      whoIsCalling: "UpdateBooking"
    };
  },

  created() {
    console.log(this.booking.time.substring(0,6))
    this.timeBegin = this.booking.time.substring(0,6 ) 
    this.timeEnd = this.booking.end.substring(0, 6)
    this.date = this.booking.date;
    this.title = this.booking.title;
    this.selectRoom = this.room.name;
  },

  methods: {
    toggleDialog() {
      if (this.dialog) {
        this.dialog = false;
        this.$emit("closeDetailCardHook");
      } else {
        this.dialog = true;
      }
    },

    closeCreateModal() {
      this.dialog = false;
      this.$emit("closeDetailCardHook");
    },
    updateBookingAjax(datetimeBegin, datetimeEnd, title, recurringType, newRoom) {
        console.log(datetimeBegin, "datetimebegin")
      this.$refs.UpdateBooking.updateBooking(
        this.booking.eventID,
        datetimeBegin,
        datetimeEnd,
        title,
        newRoom,
        this.room.id

      );
    }
  },

  computed: {
    bannerURL() {
      return require("../assets/meeting.jpg");
    }
  },

  components: {
    BookingTemplate,
    UpdateBooking
  }
};
</script>
