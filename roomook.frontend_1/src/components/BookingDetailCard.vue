<template>
<v-card hover>
      <v-img
      src="https://aws-tiqets-cdn.imgix.net/images/content/ca3a0de0cab6496b9578a2b2c300151a.jpg?auto=format&fit=crop&ixlib=python-1.1.2&q=25&s=3731f5dca39321f37cbd51be11c9abfd&w=400&h=320&dpr=2.625"
      aspect-ratio="2.75"
    >


    
            <v-layout v-if="isOwner" fill-height align-end justify-end>
  <v-btn small right fab color="orange" @click="updateBooking()">
          <v-icon>edit</v-icon>
        </v-btn>
        <v-btn small right fab color="red" @click="deleteBooking()">
          <v-icon>delete</v-icon>
        </v-btn>
            </v-layout>
          
    
    </v-img>
    <v-card-title class="detailHeadline elevation-5">
      <v-flex xs8>
        <h3 style="text-align:left">{{clickedEvent.title}}</h3>
      </v-flex>
    </v-card-title>

    <v-card-text>
      <v-container align-content-space-around align-content-space-between>
        <v-layout wrap text-xs-left>
          <v-flex md6 xs12>
            <v-icon :color="jacobColor">event</v-icon>
            {{new Date(clickedEvent.date).toString().substring(0,15)}}
          </v-flex>
          <v-flex xs12 md6>
            <v-icon :color="jacobColor">timelapse</v-icon>
            {{convertMinsToHrsMins(clickedEvent.duration)}}
          </v-flex>

          <v-flex xs12 md6>
            <v-icon :color="jacobColor">access_time</v-icon>
            {{clickedEvent.time.substring(0,6)}}
          </v-flex>
          <v-flex xs12 md6>
            <v-icon :color="jacobColor">access_time</v-icon>
            {{clickedEvent.end.substring(0,6)}}
          </v-flex>
        </v-layout>
      </v-container>
    </v-card-text>

    <v-card-actions>
      <v-flex xs6>
        <v-btn id="backButton" @click="closeDetailCardHook()" flat :color="jacobColor">Close</v-btn>
      </v-flex>
      <v-flex xs6>
        <span>{{clickedEvent.owner}}</span>
      </v-flex>
    </v-card-actions>
    <UpdateBookingDialog
      ref="updateBookingDialog"
      @closeDetailCardHook="closeDetailCardHook"
      :activeRoomID="activeRoomID"
      :booking="clickedEvent"
      :room="room"
    />
    <DeleteBookingDialog
      ref="deleteBookingDialog"
      @closeDetailCardHook="closeDetailCardHook"
      :activeRoomID="activeRoomID"
      :booking="clickedEvent"
      :room="room"
    />
</v-card>
</template>

<script>
import DeleteBookingDialog from "./DeleteBookingDialog.vue";
import UpdateBookingDialog from "./UpdateBookingDialog.vue";
export default {
  name: "BookingDetailCard",
  props: {
    clickedEvent: { type: Object },
    activeRoomID: { type: String },
    room: { type: Object },
    loggedInUser: {type: String}
  },
  data() {
    return {
      isOwner: false,
      jacobColor: "#274d6f"
    };
  },

  created(){
     if(this.clickedEvent.owner === this.loggedInUser ){
      this.isOwner = true
    } else {
      this.isOwner = false
    }
  },
  updated(){
    if(this.clickedEvent.owner === this.loggedInUser ){
      this.isOwner = true
    } else {
      this.isOwner = false
    }
  },

  methods: {

    updateBooking() {
      this.$refs.updateBookingDialog.toggleDialog();

    },
    deleteBooking() {
      this.$refs.deleteBookingDialog.toggleDialog();
    },
    convertMinsToHrsMins: minutes => {
      if (minutes <= 60) {
        return minutes + " min";
      }

      var h = Math.floor(minutes / 60);
      var m = minutes % 60;
      h = h < 10 ? "0" + h : h;
      m = m < 10 ? "0" + m : m;
      return h + "h " + m + "min";
    },
    closeDetailCardHook() {
      this.$emit("closeDetailCardHook");
    }
  },

  components: {
    DeleteBookingDialog,
    UpdateBookingDialog
  }
};
</script>

<style scope>
.detailCard {
  width: 100%;
  margin-left: 10px;
  margin-right: 10px;
  text-align: center;
}

#backButton {
  background-color: rgb(39, 77, 111, 0.2);
}

.detailHeadline {
  padding-top: 5px;
  padding-bottom: 5px;
  background-color: #274d6f;
  width: 100%;
  color: white;
}
</style>