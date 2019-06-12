<template>
  <div>
    <v-btn type="button" @click="toggle">Fullscreen</v-btn>
    <fullscreen ref="fullscreen" @change="fullscreenChange">
      <v-flex
        class="container"
        v-bind:class="{ active: isActive, 'blur': hasBlur }"
        grid-list-xl
        fluid
      >
        <v-layout v-if="isoccupied" row wrap>
          <v-layout xs6 fill-height>
            <v-flex class="text-xs-center" id="circ">
              <v-progress-circular
                :rotate="-90"
                :size="480"
                :width="6"
                :value="value"
                color="#FF1744"
              >
                <v-flex xs12 class="text uhrfrei" align-self-center>{{time | moment("HH:mm")}}</v-flex>
                <v-flex xs12 class="roomID">Aquarium</v-flex>
              </v-progress-circular>
            </v-flex>
          </v-layout>

          <v-flex xs6>
            <v-layout aglin-start justify-start column wrap>
              <v-flex class v-if="isoccupied" xs12>
                <v-card dark class="card currentBooking">
                  <v-card-text class="aktuellertermin text-xs-left">{{currentBooking.title}}</v-card-text>
                  <v-flex
                    xs12
                    class="text-xs-left currentBookingTime"
                  >now - {{currentBooking.end | moment("HH:mm")}}</v-flex>
                  <v-progress-linear :indeterminate="true" color="#FF1744" height="0.3rem"></v-progress-linear>
                  <v-flex xs12 class="text-xs-right">
                    <v-btn
                      small
                      right
                      color="#FF1744"
                      @click="dialogcheckout = true, blurTrue()"
                    >CHECKOUT</v-btn>
                    {{currentBooking.userID}}
                  </v-flex>
                </v-card>
              </v-flex>

              <v-flex v-for="meeting in sliceBooking" :key="meeting.bookingID" xs2>
                <v-card class="cardNext text-xs-left" flat dark>
                  <v-card-title>
                    <v-flex xs12 class="nextTitle">{{meeting.title}}</v-flex>
                    <v-flex
                      xs6
                    >{{meeting.begin | moment("HH:mm")}} - {{meeting.end | moment("HH:mm")}}</v-flex>

                    <v-flex xs6 class="text-xs-right">{{meeting.userID}}</v-flex>
                  </v-card-title>
                </v-card>
              </v-flex>
            </v-layout>
          </v-flex>
        </v-layout>

        <!--////////////////////////////////////////////////////////////////-->

        <!--CIRC WENN FREI / GRÜN--->

        <v-layout v-if="!isoccupied" row wrap>
          <v-flex xs6 class="margintop text-xs-center">
            <v-progress-circular :rotate="-90" :size="480" :width="6" :value="100" color="#00E676">
              <v-flex xs12 class="text uhrfrei" align-self-center>{{time | moment("HH:mm")}}</v-flex>
              <v-flex xs12 class="roomID">Aquarium</v-flex>
            </v-progress-circular>
          </v-flex>
          <v-flex class="text-xs-left" xs6>
            <v-layout aglin-start justify-start column wrap>
              <v-flex xs8 v-if="isNextBookings">
                <v-btn
                  round
                  block
                  color="#00E676"
                  class="btn"
                  dark
                  large
                  @click.stop="dialog = true"
                  @click="blurTrue()"
                >book</v-btn>
              </v-flex>
              <v-flex v-if="isNextBookings">
                <h1>upcoming</h1>
              </v-flex>

              <v-flex v-for="meeting in sliceBooking" :key="meeting.bookingID" xs2>
                <v-card class="cardNext text-xs-left" flat dark>
                  <v-card-title class="next" primary-title>
                    <v-flex xs12 class="nextTitle">{{meeting.title}}</v-flex>
                    <v-flex
                      xs6
                    >{{meeting.begin | moment("HH:mm")}} - {{meeting.end | moment("HH:mm")}}</v-flex>

                    <v-flex xs6 class="text-xs-right">{{meeting.userID}}</v-flex>
                  </v-card-title>
                </v-card>
              </v-flex>
            </v-layout>

            <v-flex xs10 class="nobookings" v-if="!isNextBookings">
              <h3>No upcoming bookings</h3>
            </v-flex>

            <v-flex xs8 v-if="!isNextBookings">
              <v-btn
                round
                block
                color="#00E676"
                class="btn"
                dark
                large
                @click.stop="dialog = true"
                @click="blurTrue()"
              >book</v-btn>
            </v-flex>
          </v-flex>
        </v-layout>

        <!--///////////BOTTOM NAV OCCUPIED//////////////////////////////////////////////////////////////////-->
        <v-bottom-nav
          value="true"
          fixed
          style="height:unset"
          class="belegt"
          color="#FF1744"
          v-if="isoccupied"
        >
          <v-layout justify-center row>
            <v-flex xs2>
              <span class="material-icons">remove_circle_outline</span>
            </v-flex>
          </v-layout>
        </v-bottom-nav>
        <!--///////////////////////////////////////////////////////////////////////////////////////////////-->

        <!--///////////BOTTOM NAV FREE///////////////////////////////////////////////////////////////////////////-->

        <v-bottom-nav
          color="#00E676"
          value="true"
          fixed
          style="height:unset"
          class="belegt"
          v-if="!isoccupied"
        >
          <v-flex v-if="!isoccupied" xs6 class="rest" align-self-center>
            <span class="material-icons">done</span>
          </v-flex>
        </v-bottom-nav>
      </v-flex>

      <v-dialog
        v-model="dialogcheckout"
        v-if="currentBooking != null"
        persistent
        lazy
        width="600px"
      >
        <v-card dark>
          <v-card-title></v-card-title>
          <v-card-text>
            <v-layout align-start justify-start row wrap>
              <v-flex class="checkout" xs12>{{currentBooking.title}}</v-flex>
              <v-flex class="checkout" xs12>als {{currentBooking.userID}} auschecken?</v-flex>
              <v-flex xs10 offset-xs1>
                <v-progress-linear color="grey" height="1" value="100"></v-progress-linear>
              </v-flex>

              <v-flex xs2 offset-xs6>
                <v-btn color="#FF1744" @click="dialogcheckout = false,checkOut()">Checkout</v-btn>
              </v-flex>
              <v-flex xs4>
                <v-btn color="grey" @click="dialogcheckout = false, bookingCancel()">Abbrechen</v-btn>
              </v-flex>
            </v-layout>
          </v-card-text>
        </v-card>
      </v-dialog>
      <!--///////////////////////////////////////////////////////////////////////////////////////-->
      <v-dialog v-model="dialog" persistent lazy width="700px" height="680px">
        <v-time-picker
          class="noblur"
          header-color="#00E676"
          color="#00E676"
          dark
          full-width
          :landscape="landscape"
          format="24hr"
          v-model="bookingtime"
          :allowed-minutes="allowedStep"
        >
          <v-spacer>Terminende festlegen</v-spacer>
          <v-layout justify-end row wrap>
            <v-flex xs4>
              <v-btn color="#FF1744" @click="dialog = false, bookingCancel()">CANCEL</v-btn>
            </v-flex>
            <v-flex xs6>
              <v-btn color="#00E676" @click="tabletBooking()">BUCHEN</v-btn>
            </v-flex>
          </v-layout>
        </v-time-picker>
        <v-alert
          :value="pickerror"
          type="error"
          color="red"
        >Bitte geben Sie die Zeit in HH:mm Format an.</v-alert>
      </v-dialog>
    </fullscreen>
  </div>
</template>

<script>
import moment from "moment";
import Fullscreen from "vue-fullscreen/src/component.vue";
import axios from "axios";
var schedule = require("node-schedule");

export default {
  data() {
    return {
      interval: null,
      value: 0,
      time: moment(),
      bookings: [],
      sliceBooking: [],
      begin: "",
      end: "",
      isoccupied: null,
      currentBooking: null,
      duration: 0,
      isNextBookings: false,
      fullscreen: false,
      landscape: true,
      picker: null,
      dialog: false,
      bookingtime: null,
      isActive: true,
      hasBlur: false,
      tabletBook: "",
      pickerror: false,
      dialogcheckout: false
    };
  },

  beforeDestroy() {
    clearInterval(this.interval);
  },

  created() {
    this.$options.sockets.onmessage = () => {
      console.log("websockets");
      axios

        .get(this.$url + "/tablet/b/r/" + this.$route.params.id)
        .then(response => {
          this.bookings = response.data;
          this.getCurrentBookings();
        })
        .catch(error => console.log(error));
    };
    this.interval = setInterval(() => {
      if (this.value > 100) {
        return (this.value = 0);
      }
      if (this.currentBooking != null) {
        this.value += 100 / (this.duration * 60);
        this.time = moment();
      }
      this.time = moment();
    }, 1000);

    this.fetchBookings();

    var func = this.fetchBookings;

    schedule.scheduleJob("0 7 * * *", function() {
      func();
    });
  },

  methods: {
    allowedStep: m => m % 5 === 0,
  

    checkOut() {
      this.hasBlur = false;
      axios
        .put("/tablet/booking/" + this.currentBooking.bookingID, {
          title: this.currentBooking.title,
          roomID: this.currentBooking.roomID,
          begin: [this.currentBooking.begin[0]],
          end: [this.time.format("YYYY-MM-DD HH:mm:ss")]
        })
        .then(response => {
          this.$socket.sendObj({
            Type: "createBooking",
            RoomID: this.$route.params.id
          });
        })
        .catch(err => {
          console.log(err);
        });
    },

    bookingCancel() {
      this.bookingtime = null;
      this.hasBlur = false;
      this.isActive = true;
    },
    tabletBooking() {
      var tabletHHmm = this.bookingtime;
      if (tabletHHmm === null) {
        this.pickerror = true;
        this.dialog = true;
      } else {
        var ymd = moment().format();
        ymd = ymd.substring(0, 10);
        this.tabletBook = ymd + " " + tabletHHmm + ":" + "00";
        var test = moment().format();
        this.pickerror = false;
        this.dialog = false;
        this.hasBlur = false;
        console.log(this.tabletBook);
        axios.post("/tablet/booking/", {
          title: "tablet booking",
          roomID: this.$route.params.id,
          begin:  [this.time.format("YYYY-MM-DD HH:mm:ss")],
          end: [this.tabletBook],
          recurringType: "single"

        }).then(response => {
          this.$socket.sendObj({
            Type: "createBooking",
            RoomID: this.$route.params.id
          }),
          this.bookingtime = null
        })
        .catch(err => {
          console.log(err);
        });
      }
    },

    blurTrue() {
      this.isActive = false;
      this.hasBlur = true;

      console.log("btn clicked");
    },

    bookingDuration() {
      var start = moment(this.currentBooking.begin[0]);
      var end = moment(this.currentBooking.end[0]);
      var ergebnis = 0;
      if (this.time != start) {
        ergebnis = moment.duration(end.diff(this.time)).asMinutes();
        this.duration = moment.duration(end.diff(start)).asMinutes();
        this.value = (this.duration - ergebnis) / (this.duration / 100);
      } else {
        ergebnis = moment.duration(end.diff(start));
        this.duration = ergebnis.asMinutes();
      }
    },

    getCurrentBookings() {
      var x = [];
      this.bookings.forEach(element => {
        if (
          moment(element.end[0]) >= moment() &&
          moment(element.end[0]).get("year") == moment().get("year") &&
          moment(element.end[0]).get("month") == moment().get("month") &&
          moment(element.end[0]).get("date") == moment().get("date")
        ) {
          x.push(element);
        }
      });
      this.bookings = x;

      this.isNextBookings = this.bookings.length != 0 ? true : false;

      if (
        this.bookings.length != 0 &&
        moment(this.bookings[0].end[0]) >= moment() &&
        moment(this.bookings[0].begin[0]) <= moment()
      ) {
        this.currentBooking = this.bookings[0];
        this.sliceBooking = this.bookings;
        this.sliceBooking = this.sliceBooking.slice(1, 3);
        this.isoccupied = true;
        this.refreshBookings();
        this.bookingDuration();
      } else {
        this.isoccupied = false;
        this.currentBooking = null;
        this.sliceBooking = this.bookings;
        this.sliceBooking = this.sliceBooking.slice(0, 3);
        this.refreshBookings();
      }
    },
    ecxec() {
      if (this.bookings.length >= 0) {
        this.getCurrentBookings();
        this.refreshBookings();
      }
    },
    refreshBookings() {
      var houres, minutes, func, timerString;
      if (this.currentBooking === null) {
        houres = this.bookings[0].begin[0].substring(11, 13);
        minutes = this.bookings[0].begin[0].substring(14, 16);
        func = this.getCurrentBookings;

        timerString = minutes + " " + houres + " * * *";

        schedule.scheduleJob(timerString, function() {
          func();
        });
      } else {
        houres = this.currentBooking.end[0].substring(11, 13);
        minutes = this.currentBooking.end[0].substring(14, 16);
        func = this.getCurrentBookings;

        timerString = minutes + " " + houres + " * * *";
        schedule.scheduleJob(timerString, function() {
          func();
        });
      }
    },
    fetchBookings() {
      axios
        .get(this.$url + "/tablet/b/r/" + this.$route.params.id)
        .then(response => {
          this.bookings = response.data;

          this.ecxec();
        });
    },
    terminDuratio() {
      this.begin = moment(this.bookings[0].begin[0]);
      this.end = moment(this.bookings[0].end[0]);
    },
    isOccupied() {
      if (this.time >= this.begin && this.time <= this.end) {
        this.isoccupied = true;
      } else {
        this.isoccupied = false;
      }
    },
    toggle() {
      this.$refs["fullscreen"].toggle();
    },
    fullscreenChange(fullscreen) {
      this.fullscreen = fullscreen;
    }
  },
  components: {
    Fullscreen
  }
};
</script>

<style scoped>
.checkout {
  font-size: 3rem;
}
.blur {
  filter: blur(7px);
}
.noblur {
  filter: blur(0px);
}
.btn {
  height: 4rem;
  font-size: 2rem;
  font-weight: 200;
}
.roomID {
  color: white;
  font-weight: 200;
  font-size: 3rem;
}
.cardnext {
  opacity: 0.6;
}

.margintop {
  margin-top: 2rem;
}
.vchipnext {
  font-size: 1rem;
  height: 3rem;
  margin-left: 1rem;
}

.vchip {
  font-size: 1.3rem;
  height: 3rem;
}
.cardtitle {
  font-size: 2rem;
}
.material-icons {
  color: aliceblue;
  font-weight: 100;
  opacity: 0.8;
  padding-top: 1rem;
  font-size: 6rem;
}
#circ {
  padding-top: -3rem;
}
.v-progress-circular {
  margin-top: 2rem;
  margin-bottom: 2rem;
}
.container {
  margin: 0vh;
  background-image: url("../assets/ttt.png");
  background-size: cover;
  height: 100vh;
  width: 100%;
  background-repeat: no-repeat;
}
.uhr {
  font-size: 7rem;
}
.rest {
  font-size: 3rem;
  font-weight: 200;
  color: #ff1744;
}
.lock {
  font-weight: 600;
  font-size: 6rem;
  padding-bottom: 1rem;
}
.text {
  color: aliceblue;
  font-weight: 50;
}
.aktuellertermin {
  font-size: 2.3rem;
  padding-top: 1rem;
  font-weight: 500;
}
.next {
  font-size: 1.5vw;
  font-weight: 300;
  align-content: center;
}
.belegt {
  color: aliceblue;
  font-size: 5rem;
  font-weight: 100;
  width: 100%;
}
.frei {
  color: aliceblue;
  font-size: 6rem;
  margin-top: 0rem;
  font-weight: 100;
}
.uhrfrei {
  font-size: 9rem;
  margin-top: 0%;
}

#roomname {
  font-size: 4rem;
  margin: auto;
  font-weight: 100;
  color: aliceblue;
}
.bookwhenfree {
  font-size: 4rem;
}
.card {
  background-color: rgb(0, 0, 0, 0.3);
  padding: 15px;
}

.cardNext {
  background-color: rgb(0, 0, 0, 0.3);
}
.currentBooking {
  box-shadow: 0px 0px 0px 1px #ff1744;
}

.currentBookingTime {
  font-size: 1.3rem;
}
.nextTitle {
  font-size: 18px;
}
h2,
h3 {
  color: rgb(255, 255, 255, 0.9);
  font-weight: 300;
}

.nobookings {
  font-size: 30px;
  margin-top: 40%;
}
</style>
