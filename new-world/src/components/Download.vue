<template>
  <v-list two-line>
    <v-row justify="space-between">
      <v-col md="4">
        <v-subheader>Download</v-subheader>
      </v-col>
      <v-col md="2" offset-md="2">
        <v-btn icon v-on:click="startDownloadItem([selected])">
          <v-icon>mdi-arrow-down-bold-circle</v-icon>
        </v-btn>
        <v-btn icon v-on:click="deleteDownloadItem([selected])">
          <v-icon>mdi-pause</v-icon>
        </v-btn>
        <v-btn icon v-on:click="pauseDownloadItem([item.id])">
          <v-icon>mdi-delete</v-icon>
        </v-btn>
      </v-col>
      <v-col>
        <download-dialog />
      </v-col>
    </v-row>

    <template v-for="(item, index) in items">
      <v-list-item :key="item.id">
        <v-list-item-action>
          <v-checkbox
            v-model="selected"
            :value="item.id"
            on-icon="mdi-circle"
            off-icon="mdi-circle-outline"
          ></v-checkbox>
        </v-list-item-action>

        <v-list-item-content>
          <v-list-item-title v-text="item.name"></v-list-item-title>

          <v-list-item-subtitle
            class="text--primary"
            v-text="resolveState(item.state)"
          ></v-list-item-subtitle>

          <v-list-item-subtitle>
            <v-progress-linear
              color="teal"
              buffer-value="0"
              :value="item.state.percent"
            ></v-progress-linear
          ></v-list-item-subtitle>
        </v-list-item-content>

        <v-list-item-action>
          <v-row>
            <v-menu offset-y>
              <template v-slot:activator="{ on, attrs }">
                <v-btn v-bind="attrs" v-on="on" icon>
                  <v-icon>mdi-dots-vertical</v-icon>
                </v-btn>
              </template>
              <v-list>
                <v-list-item>
                  <v-btn text v-on:click="startDownloadItem([item.id])"
                    >Start</v-btn
                  >
                </v-list-item>
                <v-list-item>
                  <v-btn text v-on:click="deleteDownloadItem([item.id])"
                    >Delete</v-btn
                  >
                </v-list-item>
                <v-list-item>
                  <v-btn text v-on:click="pauseDownloadItem([item.id])"
                    >Pause
                  </v-btn>
                </v-list-item>
              </v-list>
            </v-menu>
          </v-row>
        </v-list-item-action>
      </v-list-item>

      <v-divider v-if="index < items.length - 1" :key="index"></v-divider>
    </template>
  </v-list>
</template>
<script>
import DownloadDialog from "./DownloadDialog.vue";
export default {
  components: { DownloadDialog },
  name: "Download",

  data: () => ({
    sheet: false,
    progressBar: false,
    downloadDialog: false,
    newDownloadItems: "",
    selected: [],
    items: [],
  }),

  mounted() {
    this.loadData();
    const timer = setInterval(this.loadData, 5000);
    this.$once("hook:beforeDestroy", () => {
      clearInterval(timer);
    });
  },

  methods: {
    loadData() {
      this.$axios
        .get("download")
        .then((resp) => {
          this.items = resp.data;
        })
        .catch(console.error);
    },

    startDownloadItem(ids) {
      console.debug(ids);
    },

    deleteDownloadItem(ids) {
      console.debug(ids);
    },

    pauseDownloadItem(ids) {
      console.debug(ids);
    },

    resolveState(state) {
      return (
        state.percent.toFixed(2) +
        "% " +
        this.convertBytes(state.rate) +
        "/s " +
        this.convertBytes(state.complete_bytes) +
        "/" +
        this.convertBytes(state.total_bytes)
      );
    },
    convertBytes(bytes) {
      const units = ["bytes", "KB", "MB", "GB", "TB"];
      var unit = 0;
      while (bytes > 1024) {
        unit += 1;
        bytes /= 1024;
      }
      return bytes.toFixed(2) + units[unit];
    },
  },
};
</script>