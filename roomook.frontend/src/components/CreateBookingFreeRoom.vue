<template>
  <v-form ref="form">
    <v-card flat>
      <v-card-text>
        <v-layout row wrap>
          <v-flex xs12>
            <v-menu
              v-model="menuDate"
              :close-on-content-click="false"
              :nudge-right="40"
              lazy
              transition="scale-transition"
              offset-y
              full-width
              max-width="290px"
              min-width="290px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  ref="menuDate"
                  :rules="[() =>  !!date || 'Required field']"
                  v-model="date"
                  label="Choose your date"
                  persistent-hint
                  prepend-icon="event"
                  readonly
                  required
                  v-on="on"
                ></v-text-field>
              </template>
              <v-date-picker v-model="date" no-title @input="menuDate = false"></v-date-picker>
            </v-menu>
          </v-flex>
          <v-flex xs12>
            <v-layout row wrap>
              <v-flex xs12>
                <v-select
                  ref="beginhour"
                  v-model="beginhour"
                  :items="hours"
                  label="Hour"
                  prepend-icon="access_time"
                  required
                ></v-select>
              </v-flex>
              <v-flex xs12>
                <v-select
                  ref="beginmin"
                  v-model="beginmin"
                  :items="minutes"
                  label="Min"
                  prepend-icon="access_time"
                  required
                ></v-select>
              </v-flex>
            </v-layout>
          </v-flex>
          <v-flex xs12>
            <v-text-field
              ref="duration"
              mask="time"
              :rules="[duration =>  !!duration || 'Required field',
                       duration =>   ( duration.length >=4)|| 'Wrong Format',
              ]"
              v-model="duration"
              prepend-icon="access_time"
              label="Duration"
              suffix="hh:mm"
              required
            ></v-text-field>
          </v-flex>
        </v-layout>
      </v-card-text>
      <v-card-actions>
        <v-btn block flat color="blue darken-1" id="findButton" @click="SearchFreeRooms()">SEARCH</v-btn>
      </v-card-actions>
    </v-card>
  </v-form>
</template>

<script>
import axios from "axios";
import _ from "lodash";

export default {
  data() {
    return {
      freeRooms: null,
      date: new Date().toISOString().substr(0, 10),
      hours: [
        "7",
        "8",
        "9",
        "10",
        "11",
        "12",
        "13",
        "14",
        "15",
        "16",
        "17",
        "18",
        "19"
      ],
      minutes: ["00", "15", "30", "45"],
      duration: "0030",
      max: 360,

      beginmin: "",
      beginmin: [v => !!v || "Name is required"],
      beginhour: "",
      testdate: new Date(),
      menuDate: false,

      debounceFunc: _.debounce(this.fetchFreeRooms, 300),

      formHasErrors: false
    };
  },

  created() {
    this.getTimePreload();
  },

  watch: {
    date() {
      this.dateFormatted = this.formatDate(this.date);
    }
  },

  computed: {
    refreshRooms(room) {
      this.freeRooms = room;
      return this.freeRooms;
    },
    computedDateFormatted() {
      return this.formatDate(this.date);
    }
  },

  methods: {
    allowedDates: val => ![0, 6].includes(new Date(val).getDay()),

    formatDate(date) {
      if (!date) return null;

      const [year, month, day] = date.split("-");
      return `${month}/${day}/${year}`;
    },
    validate() {
      if (this.$refs.form.validate()) {
        this.snackbar = true;
      }
    },
    getTimePreload() {
      var min = new Date().getUTCMinutes();
      if (min >= 15 && min <= 29) {
        this.beginhour = (new Date().getUTCHours() + 2).toString();
        this.beginmin = "30";
      } else if (min >= 30 && min <= 44) {
        this.beginhour = (new Date().getUTCHours() + 2).toString();

        this.beginmin = "45";
      } else if (min >= 45 && min <= 59) {
        this.beginhour = (new Date().getUTCHours() + 3).toString();

        this.beginmin = "00";
      } else if (min >= 0 && min <= 14) {
        this.beginhour = (new Date().getUTCHours() + 2).toString();

        this.beginmin = "15";
      }
    },
    refreshFreeRooms(freeRooms) {
      this.$emit("refreshFreeRooms", freeRooms);
    },
    toggleRoomView(toggle) {
      this.$emit("toggleRoomView", toggle);
    },
    getFreeRooms(beginDate, endDate) {
      this.fetchFreeRoomsWithDebounce(beginDate, endDate);
    },

    SearchFreeRooms() {
      this.$refs.form.validate();
      var tmpDate = new Date(
        this.date + " " + this.beginhour + ":" + this.beginmin
      );
      tmpDate = new Date(tmpDate.setUTCHours(tmpDate.getUTCHours() + 2));
      var endTime = new Date(
        tmpDate.setMinutes(
          tmpDate.getUTCMinutes() + parseInt(this.duration.substring(2, 4), 10)
        )
      );
      if (this.duration.substring(0, 1) == "0") {
        var realEndDateTime = new Date(
          endTime.setHours(
            endTime.getHours() + parseInt(this.duration.substring(1, 2), 10)
          )
        );
      } else {
        var realEndDateTime = new Date(
          endTime.setHours(
            endTime.getHours() + parseInt(this.duration.substring(0, 2), 10)
          )
        );
      }
      var tmpBeginDate1 = new Date(
        this.date + " " + this.beginhour + ":" + this.beginmin
      );
      var realBeginDateTime = new Date(
        tmpBeginDate1.setHours(tmpBeginDate1.getHours() + 2)
      );
      this.fetchFreeRooms(realBeginDateTime, realEndDateTime);
    },

    fetchFreeRoomsWithDebounce(beginDate, endDate) {
      this.debounceFunc(beginDate, endDate);
    },

    fetchFreeRooms(beginDate, endDate) {
      axios
        .post(this.$url + "/r/d/", {
          begin: beginDate.toISOString().substr(0, 16),
          end: endDate.toISOString().substr(0, 16)
        })
        .then(response => {
          if (response.status === 401) {
            this.$router.push("/login");
          } else if (response.status === 200) {
            this.refreshFreeRooms(response.data);
            this.toggleRoomView(true);
          }
        })
        .catch(error => console.log(error));
    }
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

#findButton {
  background-color: #e4f1fc;
}
</style>