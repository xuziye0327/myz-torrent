<template>
  <v-treeview
    v-model="tree"
    :items="items"
    item-key="full_path"
    selectable
    open-on-click
    item-children="childs"
    expand-icon="mdi-chevron-down"
    on-icon="mdi-bookmark"
    off-icon="mdi-bookmark-outline"
    indeterminate-icon="mdi-bookmark-minus"
  >
    <template v-slot:prepend="{ item, open }">
      <v-icon v-if="item.is_dir">
        {{ open ? "mdi-folder-open" : "mdi-folder" }}
      </v-icon>
      <v-icon v-else> mdi-file </v-icon>
    </template>
  </v-treeview>
</template>
<script>
export default {
  name: "Files",
  data: () => ({
    tree: [],
    items: [],
  }),

  mounted() {
    this.$axios
      .get("file")
      .then((resp) => {
        this.items = resp.data;
      })
      .catch(console.error);
  },

  methods: {
    up() {
      console.log(this.tree);
    },
  },
};
</script>