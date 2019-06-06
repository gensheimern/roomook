<template></template>

<script>
import axios from "axios";
export default {
  props: {
    booking: { type: Object },
    activeRoomID: { type: String }
  },
  data() {
    return {};
  },

  methods: {
    deleteBooking(deleteRecurrent) {
      if (!deleteRecurrent) {
        axios
          .delete(this.$url + "/booking/" + this.booking.eventID)
          .then(response => {
            if (response.status === 401) {
              this.$router.push("/login");
            } else if (response.status === 200) {
              this.$emit("toggleDialog");
              this.$socket.sendObj({
                Type: "createBooking",
                RoomID: this.activeRoomID
              });
            }
          })
          .catch(err => console.log(err));
      } else {
        axios
          .delete(
            this.$url + "/booking/" + this.booking.parentID + "/recurrent/"
          )
          .then(response => {
            if (response.status === 401) {
              this.$router.push("/login");
            } else if (response.status === 200) {
              this.$emit("toggleDialog");
              this.$socket.sendObj({
                Type: "createBooking",
                RoomID: this.activeRoomID
              });
            }
          })
          .catch(err => console.log(err));
      }
    }
  }
};
</script>
