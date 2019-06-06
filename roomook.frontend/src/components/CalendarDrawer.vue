<template>
  <v-navigation-drawer class="nav_drawer" v-model="drawer" fixed app>
    <template>
      <v-toolbar flat class="transparent">
        <v-list class="pa-0">
          <v-list-tile avatar>
            <v-list-tile-avatar>
              <v-avatar size="39" :color="getAvatarColor()">
                <span class="white--text headline">{{loggedInUser.substring(0,2)}}</span>
              </v-avatar>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-layout align-center row>
                <v-flex 10 xs>
                  <v-list-tile-title>{{loggedInUser}}</v-list-tile-title>
                </v-flex>
                <v-flex offset-xs8 xs2>
                  <v-icon @click="logout()">logout</v-icon>
                </v-flex>
              </v-layout>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-toolbar>

      <v-flex class="fav-room">
        <v-flex xs10 offset-xs1>
          <h1 style="color:white; margin-bottom:20px">WHO´S YOUR FAVORITE</h1>
          <v-spacer></v-spacer>
          <FavoriteRoom ref="favoriteRoom"/>
          <v-layout align-end justify-end>
            <v-btn @click="saveFavoriteRoom()" color="green" small>
              <v-icon>save</v-icon>set favorite
            </v-btn>
          </v-layout>
        </v-flex>
      </v-flex>
      <v-layout class="drawer_category" column justify-start text-xs-left>
        <v-layout
          @click="showSearchFreeRoom = !showSearchFreeRoom"
          row
          wrap
          class="drawer_category"
        >
          <v-flex xs10>
            <h1>SEARCH FREE ROOM</h1>
          </v-flex>
          <v-flex xs1>
            <v-layout xs1 justify-center>
              <v-btn icon>
                <v-icon>{{ !showSearchFreeRoom ? 'keyboard_arrow_down' : 'keyboard_arrow_up' }}</v-icon>
              </v-btn>
            </v-layout>
          </v-flex>
        </v-layout>
        <CreateBookingFreeRoom
          v-if="showSearchFreeRoom"
          ref="BookingFreeRoom"
          @toggleRoomView="toggleRoomView"
          @refreshFreeRooms="refreshFreeRooms"
        />
        <v-divider></v-divider>
        <v-layout @click="showSelectRoom = !showSelectRoom" row wrap class="drawer_category">
          <v-flex class="drawer_category" xs10>
            <h1>SELECT ROOM</h1>
          </v-flex>
          <v-flex xs1>
            <v-layout xs1 justify-center>
              <v-btn icon>
                <v-icon>{{ !showSelectRoom ? 'keyboard_arrow_down' : 'keyboard_arrow_up' }}</v-icon>
              </v-btn>
            </v-layout>
          </v-flex>
        </v-layout>
        <v-flex v-if="showSelectRoom" offset-xs1></v-flex>
        <v-layout v-if="showSelectRoom" pa-3>
          <v-flex xs1>
            <v-treeview
              :active.sync="active"
              :items="items"
              prepend-icon="timelapse"
              :load-children="fetchRooms"
              :open.sync="open"
              activatable
              active-class="primary--text"
              open-on-click
              transition
            ></v-treeview>
          </v-flex>
        </v-layout>
        <v-divider></v-divider>
        <v-layout
          @click="showCalendarOptions = !showCalendarOptions"
          row
          wrap
          class="drawer_category"
        >
          <v-flex class="drawer_category" xs10>
            <h1>CALENDAR OPTIONS</h1>
          </v-flex>
          <v-flex xs1>
            <v-layout xs1 justify-center>
              <v-btn icon>
                <v-icon>{{ !showCalendarOptions ? 'keyboard_arrow_down' : 'keyboard_arrow_up' }}</v-icon>
              </v-btn>
            </v-layout>
          </v-flex>
        </v-layout>
        <v-flex v-if="showCalendarOptions" offset-xs1 xs10>
          <v-select
            v-model="calendarType"
            prepend-icon="timelapse"
            :items="calendarOptions"
            label="Type"
          ></v-select>

          <v-menu
            ref="startMenu"
            v-model="startMenu"
            :close-on-content-click="false"
            :nudge-right="40"
            :return-value.sync="start"
            transition="scale-transition"
            min-width="290px"
            lazy
            offset-y
            full-width
          >
            <template v-slot:activator="{ on }">
              <v-text-field
                v-model="start"
                label="Start Date"
                prepend-icon="event"
                readonly
                v-on="on"
              ></v-text-field>
            </template>
            <v-date-picker v-model="start" :allowed-dates="allowedDates" no-title scrollable>
              <v-spacer></v-spacer>
              <v-btn flat color="primary" @click="startMenu = false">Cancel</v-btn>
              <v-btn flat color="primary" @click="$refs.startMenu.save(start)">OK</v-btn>
            </v-date-picker>
          </v-menu>
        </v-flex>
      </v-layout>
      <v-divider></v-divider>
    </template>
  </v-navigation-drawer>
</template>
  </v-navigation-drawer>
</template>

<script>
const pause = ms => new Promise(resolve => setTimeout(resolve, ms));
import axios from "axios";
import FavoriteRoom from "./FavoriteRoom.vue";
import CreateBookingFreeRoom from "./CreateBookingFreeRoom.vue";
export default {
  props: {
    AllRooms: { type: Array },
    loggedInUser: { type: String }
  },
  data() {
    return {
      drawer: null,

      showSearchFreeRoom: false,
      showSelectRoom: true,
      showCalendarOptions: false,

      active: [],

      open: [],
      rooms: [],
      avatarColor: ["teal", "red", "green", "blue", "pink"],
      start: "",
      startMenu: false,
      calendarType: "week",
      calendarOptions: [
        { text: "Day", value: "day" },
        { text: "Week", value: "week" }
      ]
    };
  },
  watch: {
    selected: function(newSelected, oldSelected) {
      this.$emit("fetchBookingByRoom", newSelected);
    },
    start: function(newStart) {
      this.$emit("changeCalendarStartDate", newStart);
    },
    calendarType: function(newCalendarType, oldCalendarType) {
      this.$emit("changeCalendarType", newCalendarType);
    }
  },
  computed: {
    items() {
      return [
        {
          name: "An der Rossweid",
          children: this.rooms
        }
      ];
    },
    selected() {
      if (!this.active.length) return undefined;

      const id = this.active[0];
      return this.rooms.find(room => room.id === id);
    }
  },

  components: {
    FavoriteRoom,
    CreateBookingFreeRoom
  },

  methods: {
    async fetchRooms(item) {
      await pause(700);
      return axios
        .get(this.$url + "/room/")
        .then(json => {
          if (json.status === 401) {
            this.$router.push("/login");
          } else if (json.status === 200) {
            json.data.forEach(element => {
              item.children.push({
                id: element.roomID,
                name: element.name
              });
            });
          }
        })
        .catch(err => console.warn(err));
    },

    getAvatarColor() {
      var color = this.avatarColor[
        Math.floor(Math.random() * Math.floor(this.avatarColor.length))
      ];
      return color;
    },

    getFreeRooms(beginDate, endDate) {
      this.$refs.BookingFreeRoom.getFreeRooms(beginDate, endDate);
    },

    toggleDrawer() {
      this.drawer = !this.drawer;
    },
    refreshFreeRooms(freeRooms) {
      this.$emit("refreshFreeRooms", freeRooms);
    },
    toggleRoomView(toggle) {
      this.$emit("toggleRoomView", toggle);
    },
    saveFavoriteRoom() {
      this.$refs.favoriteRoom.saveFavoriteRoom();
    },

    allowedDates: val => ![0, 6].includes(new Date(val).getDay()),

    logout() {
      localStorage.clear();
      this.$router.push("/login");
    }
  }
};
</script>
<style scoped>
h1 {
  color: #274d6f;
}

.profile {
  cursor: pointer;
}
.fav-room {
  background-color: var(--v-primary-base);
  padding-top: 30px;
}

.drawer_category {
  padding-top: 15px;
  padding-bottom: 15px;
}

/** Hide sidebar in the nav Drawer**/
/** In Firefox**/
.nav_drawer {
  scrollbar-width: none;
}
/** In every other Browser **/
::-webkit-scrollbar {
  display: none;
}

.fade-enter-active,
.fade-leave-active {
  transition: all 0.2s ease-in-out;
}

.fade-enter,
.fade-leave-active {
  opacity: 0;
}
</style>



