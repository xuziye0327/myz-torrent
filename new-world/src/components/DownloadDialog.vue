<template>
  <v-dialog v-model="downloadDialog" max-width="600px">
    <template v-slot:activator="{ on, attrs }">
      <v-btn fab color="primary" dark v-bind="attrs" v-on="on">
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </template>
    <v-card>
      <v-card-title>
        <span class="headline">Download Urls</span>
      </v-card-title>

      <v-alert text type="success" v-if="submitSussMsg != ''">
        {{ submitSussMsg }}
      </v-alert>

      <v-alert text type="warning" v-model="emptyUrlWarning">
        Urls cannot be empty!
      </v-alert>

      <v-alert text type="error" v-if="submitErrMsg != ''">
        {{ submitErrMsg }}
      </v-alert>

      <v-card-text>
        <v-textarea
          clearable
          clear-icon="mdi-close-circle-outline"
          auto-grow
          placeholder="Support Urls: magnet"
          hint="Split mutiple urls with new line"
          v-model="downloadUrls"
        />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-progress-circular
          v-if="progressBar"
          slot="progress"
          color="blue"
          indeterminate
        />
        <v-btn v-else color="green darken-1" text @click="addToDownloadList">
          Download
        </v-btn>
        <v-spacer />
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  data: () => ({
    submitErrMsg: "",
    submitSussMsg: "",
    emptyUrlWarning: false,
    progressBar: false,
    downloadDialog: false,
    downloadUrls: "",
  }),

  methods: {
    addToDownloadList() {
      this.progressBar = false;
      this.emptyUrlWarning = false;
      this.submitErrMsg = "";
      this.submitSussMsg = "";
      if (this.downloadUrls.trim()) {
        this.progressBar = true;
        const items = this.downloadUrls
          .trim()
          .split("\n")
          .filter((s) => s);
        this.$axios
          .post("download", items)
          .then(() => {
            this.submitSussMsg = "Success add downloads: " + items.length;
          })
          .catch(() => (this.submitErrMsg = "server error!"))
          .finally(() => {
            this.progressBar = false;
          });
      } else {
        this.emptyUrlWarning = true;
      }
    },
  },
};
</script>