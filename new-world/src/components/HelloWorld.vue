<template>
  <div class="ma-2 pa-6 fill-height">
    <v-card class="fill-height">
      <v-row class="fill-height">
        <v-col cols="auto">
          <v-navigation-drawer permanent expand-on-hover>
            <v-list nav flat>
              <v-list-item-group mandatory color="indigo" v-model="model">
                <v-list-item
                  link
                  v-for="(item, index) in items"
                  :key="index"
                  @click="switchCard(index)"
                >
                  <v-list-item-icon>
                    <v-icon>{{ item.icon }}</v-icon>
                  </v-list-item-icon>
                  <v-list-item-title>{{ item.title }}</v-list-item-title>
                </v-list-item>
              </v-list-item-group>
            </v-list>
          </v-navigation-drawer>
        </v-col>
        <v-col>
          <component v-bind:is="items[model].title"></component>
        </v-col>
      </v-row>
    </v-card>
  </div>
</template>

<script>
import Download from "./Download.vue";

export default {
  name: "HelloWorld",

  components: {
    Download,
  },

  data: () => ({
    model: 0,
    items: [
      { title: "Download", icon: "mdi-cloud-download" },
      { title: "Files", icon: "mdi-folder" },
    ],
  }),

  methods: {
    switchCard(index) {
      this.model = index;
    },
  },
};
</script>
