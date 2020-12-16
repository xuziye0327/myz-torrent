<template>
  <v-list two-line>
    <v-row justify="space-between">
      <v-col>
        <v-subheader>Download</v-subheader>
      </v-col>

      <v-col lg="1">
        <v-btn icon>
          <v-icon>mdi-arrow-down-bold-circle</v-icon>
        </v-btn>
        <v-btn icon>
          <v-icon>mdi-pause</v-icon>
        </v-btn>
        <v-btn icon>
          <v-icon>mdi-delete</v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <template v-for="(item, index) in items">
      <v-list-item :key="item.id">
        <v-list-item-action>
          <v-checkbox
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
                <v-list-item> <v-btn text>Start</v-btn> </v-list-item>
                <v-list-item> <v-btn text>Delete</v-btn> </v-list-item>
                <v-list-item> <v-btn text>Pause </v-btn> </v-list-item>
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
export default {
  name: "Download",

  data: () => ({
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
    resolveState(state) {
      return (
        state.percent +
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