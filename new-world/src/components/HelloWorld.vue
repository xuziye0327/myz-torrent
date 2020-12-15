<template>
  <div class="ma-2 pa-6 fill-height">
    <v-card class="fill-height">
      <v-row class="fill-height">
        <v-col cols="auto">
          <v-navigation-drawer permanent expand-on-hover>
            <v-list nav>
              <v-list-item-group
                v-model="model"
                active-class="deep-purple--text text--accent-4"
              >
                <v-list-item link>
                  <v-list-item-icon>
                    <v-icon>mdi-cloud-download</v-icon>
                  </v-list-item-icon>
                  <v-list-item-title>Download</v-list-item-title>
                </v-list-item>

                <v-list-item link>
                  <v-list-item-icon>
                    <v-icon>mdi-folder</v-icon>
                  </v-list-item-icon>
                  <v-list-item-title>Files</v-list-item-title>
                </v-list-item>
              </v-list-item-group>
            </v-list>
          </v-navigation-drawer>
        </v-col>
        <v-col>
          <v-main>
            <Download />
          </v-main>
        </v-col>
      </v-row>
    </v-card>
  </div>
</template>

<script>
import "./Download";

export default {
  name: "HelloWorld",

  data: () => ({
    items: [
      {
        id: "9638456f4262c972bfccf349c2468e862db30867",
        name: "Lucifer.S02.1080p.AMZN.WEBRip.DDP5.1.x265-SiGMA[rartv]",
        state: {
          running_state: 3,
          total_bytes: 50483072494,
          complete_bytes: 50483072494,
          rate: 8358315.600359916,
          percent: 100,
        },
      },
      {
        id: "9638456f4262c972bfccf349c2468e862db30867",
        name: "Lucifer.S02.1080p.AMZN.WEBRip.DDP5.1.x265-SiGMA[rartv]",
        state: {
          running_state: 3,
          total_bytes: 50483072494,
          complete_bytes: 50483072494,
          rate: 8358315.600359916,
          percent: 100,
        },
      },
    ],
  }),
};
</script>
