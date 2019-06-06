<template>
  <v-content>
    <v-container  fluid>
      <v-layout column>
        <v-flex xs12>
          <v-flex offset-md1 md10>
            <v-range-slider ticks v-model="timeRange" :min="timeMin" :max="timeMax" :step="5"></v-range-slider>
          </v-flex>
        </v-flex>
        <v-flex md12>
          <v-layout row wrap>
            <v-flex md2>
              <h2>Begin</h2>
            </v-flex>
            <v-flex md8>
              <h2>{{ parseTimeBegin()}} - {{praseTimeEnd() }}</h2>
            </v-flex>
            <v-flex md2>
              <h2>End</h2>
            </v-flex>
          </v-layout>
        </v-flex>

        <v-flex xs12>
          <v-layout wrap align-center align-start row fill-height>
            <v-fab-transition>
              <v-btn
                id="roomBackButton"
                @click="toggleRoomView(false)"
                color="blue"
                fab
                dark
                absolute
                top
                left
              >
                <v-icon>arrow_back</v-icon>
              </v-btn>
            </v-fab-transition>
            <v-flex v-if="freeRooms.rooms === null">
              <h1 class="noEntries">No free rooms available </h1>
            </v-flex>
            <v-flex v-if="freeRooms.rooms != null">
            <v-flex md3 xs12 v-for="(n, index) in freeRooms.rooms" :key="n.roomID">
              <v-form v-model="valid" ref="forms">
                <v-spacer class="pa-2">
                  <v-card>
                    <v-img
                      src="https://cdn.vuetifyjs.com/images/cards/desert.jpg"
                      aspect-ratio="2.75"
                    ></v-img>

                    <v-card-title primary-title>
                      <h3 class="headline mb-0">{{n.name}}</h3>
                    </v-card-title>
                    <v-flex xs10 style="padding-left:16px;">
                      <v-text-field
                        clearable
                        prepend-icon="title"
                        v-model="title[index]"
                        required
                        :rules="titleRules"
                        label="Set Title"
                        maxlength="30"
                        counter
                      ></v-text-field>
                      <v-layout justify-start>
                        <v-icon>access_time</v-icon>
                        {{parseTimeBegin()}} - {{praseTimeEnd()}}
                      </v-layout>
                    </v-flex>
                    <v-layout justify-end>
                      <v-card-actions>
                        <v-btn @click="createBooking(n, index)" flat color="blue">BOOK</v-btn>
                      </v-card-actions>
                    </v-layout>
                  </v-card>
                </v-spacer>
              </v-form>
            </v-flex>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-layout>
    </v-container>
  </v-content>
</template>

<script>
import axios from "axios";
const pause = ms => new Promise(resolve => setTimeout(resolve, ms));

export default {
  props: {
    freeRooms: { type: Object }
  },
  data() {
    return {
      name: "FreeRoomWrapper",
      valid: false,
      timeRange: [
        new Date(this.freeRooms.begin).getHours() * 60 +
          new Date(this.freeRooms.begin).getMinutes(),
        new Date(this.freeRooms.end).getHours() * 60 +
          new Date(this.freeRooms.end).getMinutes()
      ],
      timeMin:
        new Date(this.freeRooms.begin).getHours() * 60 +
        new Date(this.freeRooms.begin).getMinutes() -
        15,
      timeMax:
        new Date(this.freeRooms.end).getHours() * 60 +
        new Date(this.freeRooms.end).getMinutes() +
        15,
      title: [],

      titleRules: [v => !!v || "This field is required"],
      testRooms: [
        {
          RoomID: "1",
          Name: "Aquarium"
        },
        {
          RoomID: "2",
          Name: "Lager"
        },
        {
          RoomID: "13",
          Name: "Küche"
        },
        {
          RoomID: "143",
          Name: "44B"
        },
        {
          RoomID: "ff",
          Name: "44333B"
        }
      ]
    };
  },

  watch: {
    timeRange: function(newTimeRange, oldTimerange) {
      var tmpDate1 = new Date(
        this.freeRooms.begin.substring(0, 10) +
          " " +
          this.parseTimeBegin(newTimeRange[0])
      );
      var newDateBegin = new Date(tmpDate1.setHours(tmpDate1.getHours() + 2));

      var tmpDate2 = new Date(
        this.freeRooms.end.substring(0, 10) +
          " " +
          this.praseTimeEnd(newTimeRange[1])
      );
      var newDateEnd = new Date(tmpDate2.setHours(tmpDate2.getHours() + 2));

      this.$emit("getFreeRooms", newDateBegin, newDateEnd);
    }
  },

  methods: {
    parseTimeBegin() {
      var begin;
      var beginFirst = parseInt(this.timeRange[0] / 60);
      var beginSecond = this.timeRange[0] % 60;

      if (
        beginFirst.toString().length <= 1 ||
        beginSecond.toString().length <= 1
      ) {
        if (beginFirst.toString().length <= 1) {
          var beginFirst = "0" + parseInt(this.timeRange[0] / 60);
        }
        if (beginSecond.toString().length <= 1) {
          beginSecond = "0" + (this.timeRange[0] % 60);
        }

        begin = beginFirst + ":" + beginSecond;
      } else {
        begin =
          parseInt(this.timeRange[0] / 60) + ":" + (this.timeRange[0] % 60);
      }
      return begin;
    },

    praseTimeEnd() {
      var end;
      var endFirst = parseInt(this.timeRange[1] / 60);
      var endSecond = this.timeRange[1] % 60;
      if (endFirst.toString().length <= 1 || endSecond.toString().length <= 1) {
        if (endFirst.toString().length <= 1) {
          endFirst = "0" + parseInt(this.timeRange[1] / 60);
        }
        if (endSecond.toString().length <= 1) {
          endSecond = "0" + (this.timeRange[1] % 60);
        }
        end = endFirst + ":" + endSecond;
      } else {
        end = parseInt(this.timeRange[1] / 60) + ":" + (this.timeRange[1] % 60);
      }

      return end;
    },

    refreshFreeRooms(freeRooms) {
      this.$emit("refreshFreeRooms", freeRooms);
    },

    toggleRoomView(toggle) {
      this.$emit("toggleRoomView", toggle);
    },

    createBooking(event, index) {
      if (this.$refs.forms[index].validate()) {
        axios
          .post(this.$url + "/booking/", {
            Title: this.title[index],
            UserID: "nico",
            RoomID: event.roomID,
          RecurringType: 'single',
            Begin: [
              this.freeRooms.begin.substring(0, 10) +
                " " +
                this.parseTimeBegin()
            ],
            End: [
              this.freeRooms.end.substring(0, 10) + " " + this.praseTimeEnd()
            ]
          })
          .then(async response => {
            if (response.status === 401) {
              this.$router.push("/login");
            } else if (response.status === 201) {
              this.$emit("fetchBookingByRoom", {
                id: event.roomID,
                name: event.name
              });
              await pause(1000);

              this.$socket.sendObj({
                Type: "createBooking",
                RoomID: event.roomID
              });
              this.toggleRoomView(false);
            }
          })
          .catch(error => {
            console.log(error);
          });
      }
    }
  }
};
</script>

<style scoped>
#roomBackButton {
  margin-top: 5vh;
}
.noEntries {
  margin-top:20vh;
  color: #274d6f
}
h2 {
  color: #274d6f;
}
</style>