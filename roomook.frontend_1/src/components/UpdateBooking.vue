<template></template>

<script>
import axios from "axios";
const pause = ms => new Promise(resolve => setTimeout(resolve, ms));

export default {
  data() {
    return {};
  },

  created() {},

  methods: {
    updateBooking(eventID, timeBegin, timeEnd, title, newRoom, oldRoom) {
      axios
        .put(this.$url + "/booking/" + eventID, {
          Title: title,
          RoomID: newRoom,
          Begin: timeBegin,
          End: timeEnd,
          BookingID: eventID
        })
        .then(async response => {
          if(response.status === 200) {
          this.$emit("closeCreateModal");
          this.$socket.sendObj({
            Type: "createBooking",
            RoomID: newRoom
          });
                     await  pause(1000);

                    this.$socket.sendObj({
            Type: "createBooking",
            RoomID: oldRoom
          });
          }
        })
        .catch(err => {
          if(err.status === 401) {
            this.$router.push('/login')
          }
          console.log(err)
          });
    }
  }
};
</script>
