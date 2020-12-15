<template>
  <v-list two-line>
    <v-row justify="space-between">
      <v-col>
        <v-header>Download</v-header>
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
            <v-menu offset-y>
              <template v-slot:activator="{ on, attrs }">
                <v-btn v-bind="attrs" v-on="on" icon>
                  <v-icon>mdi-dots-vertical</v-icon>
                </v-btn>
              </template>
              <v-list>
                <v-list-item> <v-btn text>Start</v-btn> </v-list-item>
                <v-list-item> <v-btn text>Delete</v-btn> </v-list-item>
                <v-list-item> <v-btn text>Pause </v-btn></v-list-item>
              </v-list>
            </v-menu>
            <!-- <v-btn icon>
              <v-icon>mdi-arrow-down-bold-circle</v-icon>
            </v-btn>

            <v-btn icon>
              <v-icon>mdi-pause</v-icon>
            </v-btn>

            <v-btn icon>
              <v-icon>mdi-delete</v-icon>
            </v-btn> -->
          </row>
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
        id: "9638456f4262c972bfccf349c2468e862db30137",
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