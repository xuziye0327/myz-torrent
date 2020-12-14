<template>
  <div class="ma-2 pa-6 fill-height">
    <v-card class="fill-height">
      <v-row class="fill-height">
        <v-col cols="auto">
          <v-navigation-drawer permanent expand-on-hover>
            <v-list nav>
              <v-list-item link>
                <v-list-item-icon>
                  <v-icon>mdi-folder</v-icon>
                </v-list-item-icon>
                <v-list-item-title>My Files</v-list-item-title>
              </v-list-item>
              <v-list-item link>
                <v-list-item-icon>
                  <v-icon>mdi-account-multiple</v-icon>
                </v-list-item-icon>
                <v-list-item-title>Shared with me</v-list-item-title>
              </v-list-item>
              <v-list-item link>
                <v-list-item-icon>
                  <v-icon>mdi-star</v-icon>
                </v-list-item-icon>
                <v-list-item-title>Starred</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-navigation-drawer>
        </v-col>
        <v-col
          ><v-list two-line>
            <v-subheader>Download</v-subheader>
            <template v-for="(item, index) in items">
              <v-list-item :key="item.name">
                <v-list-item-action>
                  <v-checkbox
                    :input-value="active"
                    on-icon="mdi-circle"
                    off-icon="mdi-circle-outline"
                  ></v-checkbox>
                </v-list-item-action>

                <v-list-item-content>
                  <v-list-item-title v-text="item.name"></v-list-item-title>

                  <v-list-item-subtitle
                    class="text--primary"
                    v-text="item.name"
                  ></v-list-item-subtitle>

                  <v-list-item-subtitle>
                    <v-progress-linear
                      color="teal"
                      buffer-value="0"
                      value="20"
                      stream
                    ></v-progress-linear
                  ></v-list-item-subtitle>
                </v-list-item-content>

                <v-list-item-action>
                  <row>
                    <v-btn icon>
                      <v-icon>mdi-arrow-down-bold-circle</v-icon>
                    </v-btn>

                    <v-btn icon>
                      <v-icon>mdi-pause</v-icon>
                    </v-btn>

                    <v-btn icon>
                      <v-icon>mdi-delete</v-icon>
                    </v-btn>
                  </row>
                </v-list-item-action>
              </v-list-item>

              <v-divider
                v-if="index < items.length - 1"
                :key="index"
              ></v-divider>
            </template>
          </v-list>
        </v-col>
      </v-row>
    </v-card>
  </div>
</template>

<script>
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
