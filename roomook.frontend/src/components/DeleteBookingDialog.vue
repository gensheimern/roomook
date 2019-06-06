<template>
  <v-layout row justify-center>
    <v-dialog v-model="dialog" max-width="500">
      <v-card>
        <v-card-title class="detailHeadline elevation-5 headline">Attention</v-card-title>

        <v-card-text class="text-xs-left">
          <v-flex xs10 offset-xs1>
            <v-flex style="margin-bottom:20px;">
              <h3>You are going to delete this Booking. Are you sure?</h3>
            </v-flex>
            <v-layout column wrap>
              <v-flex class="details">
                <span style="color:grey">Title</span>
              </v-flex>
              <v-flex>
                <h4>{{booking.title}}</h4>
              </v-flex>
              <v-layout row wrap>
                <v-flex class="details">
                  <span style="color:grey">Date</span>

                  <v-flex>
                    <h4>{{booking.date}}</h4>
                  </v-flex>
                </v-flex>
                <v-flex class="details">
                  <span style="color:grey">Time</span>

                  <v-flex>
                    <h4>{{booking.time.substring(0,6)}}</h4>
                  </v-flex>
                </v-flex>
              </v-layout>
            </v-layout>
            <v-flex v-if="isRecurring" xs12 class="recurrent">
              This booking is a recurring booking.
              <v-layout justify-start row wrap>
                <v-flex>
                  <v-switch xs4 v-model="switchRecurrent" label="Do you want to delete all?"></v-switch>
                </v-flex>
              </v-layout>
            </v-flex>
          </v-flex>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn color="blue" flat="flat" @click="dialog = false">Cancel</v-btn>

          <v-btn color="red" flat="flat" @click="deleteBooking()">{{deleteButtonText}}</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <DeleteBooking
      ref="deleteBooking"
      @toggleDialog="toggleDialog"
      :booking="booking"
      :activeRoomID="activeRoomID"
    />
  </v-layout>
</template>

<script>
import DeleteBooking from "./DeleteBooking.vue";
export default {
  props: {
    booking: { type: Object },
    activeRoomID: { type: String }
  },
  data() {
    return {
      dialog: false,
      isRecurring: null,
      switchRecurrent: false,
      deleteButtonText: "DELETE"
    };
  },

  components: {
    DeleteBooking
  },

  created() {
    if (this.booking.recurringType == "single") {
      this.isRecurring = false;
    } else {
      this.isRecurring = true;
    }
  },

  watch: {
    switchRecurrent: function(newSwitchRecurrent) {
      if (newSwitchRecurrent) {
        this.deleteButtonText = "DELETE ALL";
      } else {
        this.deleteButtonText = "DELETE";
      }
    }
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

    deleteBooking() {
      this.$refs.deleteBooking.deleteBooking(this.switchRecurrent);
    }
  }
};
</script>

<style scoped>
.headline {
  padding-top: 5px;
  padding-bottom: 5px;
  background-color: #274d6f;
  width: 100%;
  color: white;
}
.details {
  margin-top: 5px;
}

.recurrent {
  margin-top: 15px;
}
</style>