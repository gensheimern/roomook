<template>
  <v-card flat>
    <v-form v-model="valid" ref="form">
      <v-toolbar dark color="#274d6f">
        <v-toolbar-title>Room selection</v-toolbar-title>
        <v-autocomplete
          clearable
          v-model="selectRoomModel"
          :items="items"
          :loading="isLoading"
          :search-input.sync="search"
          cache-items
          class="mx-3"
          hide-no-data
          hide-selected
          hide-details
          item-text="Description"
          item-value="API"
          label="Start typing to find your room"
          solo-inverted
        ></v-autocomplete>
      </v-toolbar>
      <v-card-text>
        <v-layout row wrap>
          <v-flex offset-md1 md5 xs12>
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
                  v-model="computedDateFormatted"
                  label="Choose your date"
                  persistent-hint
                  prepend-icon="event"
                  readonly
                  v-on="on"
                  required
                  :rules="dateRules"
                ></v-text-field>
              </template>
              <v-date-picker
                v-model="dateModel"
                :allowed-dates="allowedDates"
                no-title
                @input="menuDate = false"
              ></v-date-picker>
            </v-menu>
          </v-flex>

          <v-flex md5 xs12>
            <v-text-field
              clearable
              prepend-icon="title"
              v-model="titleModel"
              label="Title"
              required
              :rules="titleRules"
              maxlength="30"
              counter
            ></v-text-field>
          </v-flex>

          <v-flex offset-md1 md5 xs12>
            <v-dialog
              ref="dialogTimeBegin"
              v-model="modalDateBegin"
              :return-value.sync="timeBegin"
              persistent
              lazy
              full-width
              width="290px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="timeBegin"
                  label="Begin"
                  prepend-icon="access_time"
                  readonly
                  v-on="on"
                  required
                  :rules="beginRules"
                ></v-text-field>
              </template>
              <v-time-picker
                color="#274d6f"
                v-if="modalDateBegin"
                v-model="timeBegin"
                :scrollable="true"
                full-width
                format="24hr"
                min="07:00"
                max="19:00"
              >
                <v-spacer></v-spacer>
                <v-btn flat color="primary" @click="modalDateBegin = false">Cancel</v-btn>
                <v-btn flat color="primary" @click="$refs.dialogTimeBegin.save(timeBegin)">OK</v-btn>
              </v-time-picker>
            </v-dialog>
          </v-flex>
          <v-flex md5 xs12>
            <v-dialog
              ref="dialogTimeEnd"
              v-model="modalTimeEnd"
              :return-value.sync="timeEnd"
              persistent
              lazy
              full-width
              width="290px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="timeEnd"
                  label="End"
                  prepend-icon="access_time"
                  readonly
                  v-on="on"
                  required
                  :rules="endRules"
                ></v-text-field>
              </template>
              <v-time-picker
                v-if="modalTimeEnd"
                v-model="timeEnd"
                full-width
                color="#274d6f"
                :scrollable="true"
                format="24hr"
                min="07:00"
                max="19:00"
              >
                <v-spacer></v-spacer>
                <v-btn flat color="primary" @click="modalTimeEnd = false">Cancel</v-btn>
                <v-btn flat color="primary" @click="$refs.dialogTimeEnd.save(timeEnd)">OK</v-btn>
              </v-time-picker>
            </v-dialog>
          </v-flex>
          <v-flex v-if="showRecurrently" offset-md1 md5 xs12>
            <v-switch v-model="switchRecurrently" label="Recurrently"></v-switch>
          </v-flex>
          <v-flex v-if="showRecurrently" md5 xs12>
            <v-select
              prepend-icon="access_time"
              v-model="selectedPeriod"
              :items="periodOfTime"
              label="Period of Time"
              :disabled="!switchRecurrently"
            ></v-select>
          </v-flex>
        </v-layout>
        <v-alert v-if="createAlert" :value="true" :color="alertColor">{{alertResponse}}</v-alert>
      </v-card-text>
      <v-card-actions>
        <v-btn block color="red" flat @click="closeCreateModal()">Close</v-btn>

        <v-btn
          block
          color="blue darken-1"
          flat
          @click="callParentsAjaxFunction()"
        >{{getAcceptMessage()}}</v-btn>
      </v-card-actions>
    </v-form>
  </v-card>
</template>

<script>
import moment from "moment";
import axios from "axios";

// Constants which defines the maxium lenght of recurrent bookings
// Day = 182; 2Day = 91; Week=26, Month = 6 -> create recurrent bookings for half a year
const DAY_MAX_ITERATION = 182
const TWO_DAY_MAX_ITERATION = 91
const WEEKLY_MAX_ITERATION = 26
const MONTH_MAX_ITERATION = 6


export default {
  props: {
    timeBegin: { type: String },
    timeEnd: { type: String },
    title: { type: String },
    date: { type: String },
    selectRoom: { type: String },
    showRecurrently: { type: Boolean },
    whoIsCalling: { type: String }
  },
  data: vm => {
    return {
      valid: false,

      titleModel: "",

      entries: [],
      isLoading: false,
      search: "*",
      searchObj: null,
      selectRoomRules: [v => !!v || "This field is required"],

      selectRoomModel: "",

      titleRules: [v => !!v || "This field is required"],
      descriptionLimit: 60,

      switchRecurrently: false,
      selectedPeriod: "",
      periodOfTime: ["every day", "2-days", "weekly", "per month"],
      recurringType: "",

      menuDate: false,
      dateModel: new Date().toISOString().substr(0, 10),
      dateFormatted: vm.formatDate(new Date().toISOString().substr(0, 10)),

      dateRules: [v => !!v || "This field is required"],

      menuTimeBegin: false,
      modalDateBegin: false,
      beginRules: [v => !!v || "This field is required"],

      endRules: [v => !!v || "This field is required"],
      modalTimeEnd: false,

      createAlert: false,
      alertResponse: "",
      alertColor: "",

      room: null,
      booking: []
    };
  },

  created() {
    this.dateModel = this.date;
    this.titleModel = this.title;
    this.selectRoomModel = this.selectRoom;
  },

  computed: {
    computedDateFormatted() {
      return this.formatDate(this.dateModel);
    },

    items() {
      return this.entries.map(entry => {
        const Description = entry.name;

        return Object.assign({}, entry, { Description });
      });
    }
  },

  watch: {
    room: function(newRoom) {
      if (this.room != null) {
        this.selectRoomModel = this.room.name;
      }
    },
    date(val) {
      this.dateFormatted = this.formatDate(this.dateModel);
    },
    search(val) {
      // Items have already been loaded
      if (this.items.length > 0) return;

      // Items have already been requested
      if (this.isLoading) return;

      this.isLoading = true;

      axios
        .get(this.$url + "/room/")
        .then(res => {
          if (res.status === 401) {
            this.$router.push("/login");
          } else if (res.status == 200) {
            this.count = 33;
            res.data.forEach(element => {
              this.entries.push({
                id: element.roomID,
                name: element.name
              });
            });
          }
        })
        .catch(err => {
          console.log(err);
        })
        .finally(() => (this.isLoading = false));
    }
  },

  methods: {
    alert(createAlert, alertColor, alertResponse) {
      this.createAlert = createAlert;
      this.alertResponse = alertResponse;
      this.alertColor = alertColor;
    },
    callParentsAjaxFunction() {
      if (this.getSearchObject() === undefined) {
        this.alertColor = "error";
        this.createAlert = true;

        this.alertResponse = "Please select a room";
        this.$refs.form.validate();
      } else if (!this.$refs.form.validate()) {
        this.$refs.form.validate();
      } else {
        var dateTimes = this.createDates();
        var datetimeBegin = dateTimes[0];
        var datetimeEnd = dateTimes[1];
        //Call parents ajax function
        this.$emit(
          "AjaxRequest",
          datetimeBegin,
          datetimeEnd,
          this.titleModel,
          this.recurringType,
          this.getSearchObject().id
        );
      }
    },
    getSearchObject() {
      return this.entries.find(search => search.name === this.search);
    },
    formatDate(dateModel) {
      if (!dateModel) return null;

      const [year, month, day] = dateModel.split("-");
      return `${month}/${day}/${year}`;
    },
    parseDate(dateModel) {
      if (!dateModel) return null;

      const [month, day, year] = dateModel.split("/");
      return `${year}-${month.padStart(2, "0")}-${day.padStart(2, "0")}`;
    },
    allowedDates: val => ![0, 6].includes(new Date(val).getDay()),
    setRoom(room) {
      this.room = room;
    },

    getAcceptMessage() {
      if (this.whoIsCalling === "UpdateBooking") {
        return "UPDATE";
      }

      if (this.whoIsCalling === "CreateBooking") {
        return "CREATE";
      }
    },

    createDates() {
      var datetimeBegin = [];
      var datetimeEnd = [];
      var tmpDate = new Date(this.dateModel);

      if (this.switchRecurrently == true) {
        datetimeBegin[0] = this.dateModel + " " + this.timeBegin + ":00";
        datetimeEnd[0] = this.dateModel + " " + this.timeEnd + ":00";

        switch (this.selectedPeriod) {
          case "every day":
            this.recurringType = "every day";
            var i = 1
            while ( i <= DAY_MAX_ITERATION ) {
              if (tmpDate.getDay() != 0 && tmpDate.getDay() != 6) {
                datetimeBegin[i] =
                  tmpDate.toISOString().substr(0, 10) + " " + this.timeBegin + ":00";
                datetimeEnd[i] =
                  tmpDate.toISOString().substr(0, 10) + " " + this.timeEnd + ":00";

                tmpDate.setDate(tmpDate.getDate() + 1);
                i++;
              } else {
                tmpDate = new Date(this.getNextWeekday(tmpDate));
              }
            }
            break;

          case "2-days":
            this.recurringType = "2-days"
            var i = 1
            while ( i <= TWO_DAY_MAX_ITERATION ) {
              if (tmpDate.getDay() != 0 && tmpDate.getDay() != 6) {
                datetimeBegin[i] =
                  tmpDate.toISOString().substr(0, 10) + " " + this.timeBegin + ":00";
                datetimeEnd[i] =
                  tmpDate.toISOString().substr(0, 10) + " " + this.timeEnd + ":00";
                tmpDate.setDate(tmpDate.getDate() + 2);
                i++;
              } else {
                tmpDate = new Date(this.getNextWeekday(tmpDate));
              }
            }

            break;

          case "weekly":
            this.recurringType = "weekly";
            var i = 1
            while ( i <= WEEKLY_MAX_ITERATION ) {
              if (tmpDate.getDay() != 0 && tmpDate.getDay() != 6) {
                datetimeBegin[i] =
                  tmpDate.toISOString().substr(0, 10) + " " + this.timeBegin + ":00";
                datetimeEnd[i] =
                  tmpDate.toISOString().substr(0, 10) + " " + this.timeEnd + ":00";
                tmpDate.setDate(tmpDate.getDate() + 7);

                i++;
              } else {
                tmpDate = new Date(this.getNextWeekday(tmpDate));
              }
            }
            break;

          case "per month":
            this.recurringType = "per month";
            var i = 1
            while (i <= MONTH_MAX_ITERATION ) {
              if (tmpDate.getDay() != 0 && tmpDate.getDay() != 6) {
                datetimeBegin[i] =
                  tmpDate.toISOString().substr(0, 10) + " " + this.timeBegin + ":00";
                datetimeEnd[i] =
                  tmpDate.toISOString().substr(0, 10) + " " + this.timeEnd + ":00";
                tmpDate.setMonth(tmpDate.getMonth() + 1);
                i++;
              } else {
                tmpDate = new Date(this.getNextWeekday(tmpDate));
              }
            }
            break;
        }
      } else {
        this.recurringType = "single";
        datetimeBegin[0] = this.dateModel + " " + this.timeBegin + ":00";
        datetimeEnd[0] = this.dateModel + " " + this.timeEnd + ":00";
      }
      return [datetimeBegin, datetimeEnd];
    },

    getNextWeekday(date) {
      var tomorrow = new Date(date.setDate(date.getDate() + 1));
      return tomorrow.getDay() % 6 ? tomorrow : this.getNextWeekday(tomorrow);
    },
    closeCreateModal() {
      this.$emit("closeCreateModal");
    }
  }
};
</script>
