<template>
                <v-autocomplete
                
                  clearable
                  v-model="model"
                  :items="items"
                  :loading="isLoading"
                  :search-input.sync="search"
                  color="primary"
                  hide-no-data
                  hide-selected
                  item-text="Description"
                  item-value="API"
                  placeholder="Favorite Room"
                  label="fav room"
                  return-object
                  solo-inverted
                  dark
                ></v-autocomplete>
</template>
<script>
import axios from "axios";
export default {

  data() {
    return {
      label: null,
      search: "*",
      searchObj: null,

      descriptionLimit: 60,
      entries: [],
      isLoading: false,
      model: null
    };
  },

  created(){
    if(this.$cookies.get("favoriteRoom")!= undefined && this.$cookies.get("favoriteRoom") != null){
      this.model = this.$cookies.get("favoriteRoom").name
      this.label = 'Favorite Room'
    }
  },

  computed: {

    fields() {
      if (!this.model) return [];

      return Object.keys(this.model).map(key => {
        return {
          key,
          value: this.model[key] || "n/a"
        };
      });
    },
    items() {
      return this.entries.map(entry => {
        const Description = entry.name;

        return Object.assign({}, entry, { Description });
      });
    }
  },
  watch: {
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
          } else if (res.status === 200) {
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
    getSearchObject() {
      return this.entries.find(search => search.name === this.search);
    },
    saveFavoriteRoom() {
      if (this.search != null) {
        this.$cookies.set("favoriteRoom", this.getSearchObject());
      }
      this.dialog = false;
    }
  }
};
</script>

  