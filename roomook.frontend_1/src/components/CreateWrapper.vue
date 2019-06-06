<template>
  <v-dialog
    v-model="dialog"
    width="700"
    scrollable
    transition="dialog-bottom-transition"
    hide-overlay
  >
    <template v-slot:activator="{ on }"></template>

    <v-card >
      <v-img class="createCardImage" :src="bannerURL" aspect-ratio="1.9"></v-img>

      <v-layout row wrap>
        <v-flex xs12>
          <CreateBooking 
            ref="createBooking"
            @toggleModal="toggleModal"
            @loadingWhileCreate="loadingWhileCreate"
          />
        </v-flex>
      </v-layout>
    </v-card>
  </v-dialog>
</template>

<script>
import CreateBooking from "./CreateBooking.vue";

export default {
  data() {
    return {
      dialog: false
    };
  },
  components: {
    CreateBooking
  },

  computed: {
    bannerURL() {
      return require("../assets/meeting.jpg");
    }
  },

  methods: {
    toggleModal() {
      if (this.dialog == true) {
        this.dialog = false;
      } else {
        this.$refs.createBooking.UpdateTime()
        this.dialog = true;
      }
    },
    loadingWhileCreate(booking) {
      this.$emit("loadingWhileCreateWrapper", booking);
    },

    setRoom(room) {
       this.$refs.createBooking.setRoom(room);
    }
  }
};
</script>

<style >
.v-window__container {
  margin-top: 70px !important;
}
</style>

