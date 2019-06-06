<template>
  <v-toolbar fixed app dark color="#274d6f">
    <v-toolbar-side-icon @click.stop="toggleDrawer()" font-color="white"></v-toolbar-side-icon>
    <v-flex>
      <h1>Search Free Room</h1>
    </v-flex>
  </v-toolbar>
</template>

<script>
export default {
  name: "HeaderFindRoom",
  data() {
    return {};
  },

  methods: {
    toggleDrawer() {
      this.$emit("toggleDrawer");
    }
  }
};
</script>
