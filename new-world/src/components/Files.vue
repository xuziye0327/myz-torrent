<template>
  <v-card flat>
    <v-alert v-if="copySuccessed" dense text type="success">
      {{ copySuccessed }}
    </v-alert>
    <v-alert v-if="copyFailed" dense type="error"> </v-alert>
    <v-footer padless>
      <v-card flat tile width="100%">
        <v-card-text>
          <v-btn icon @click="copyFileLinks">
            <v-icon>mdi-content-copy</v-icon>
          </v-btn>
          <v-btn icon @click="deleteFiles">
            <v-icon>mdi-delete-outline</v-icon>
          </v-btn>
        </v-card-text>
      </v-card>
    </v-footer>
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
        <v-icon v-if="item.childs">
          {{ open ? "mdi-folder-open" : "mdi-folder" }}
        </v-icon>
        <v-icon v-else> mdi-file </v-icon>
      </template>
    </v-treeview>
  </v-card>
</template>
<script>
export default {
  name: "Files",
  data: () => ({
    tree: [],
    items: [],
    copySuccessed: "",
    copyFailed: "",
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
    copyFileLinks() {
      this.copySuccessed = this.copyFailed = "";
      const msg = this.tree
        .map(
          (link) =>
            process.env.VUE_APP_DEV_SERVER + "file/" + this.pathEncode(link)
        )
        .join("\n");
      this.$copyText(msg).then(
        (e) => {
          this.copySuccessed = "Copied!";
          console.log(e);
        },
        (e) => {
          this.copyFailed = e.text;
          console.error(e);
        }
      );
    },

    deleteFiles() {},

    pathEncode(path) {
      path = encodeURI(path);
      path = window.btoa(path);
      return path;
    },
  },
};
</script>