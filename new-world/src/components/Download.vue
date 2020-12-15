<template>
  <v-list two-line>
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

      <v-divider v-if="index < items.length - 1" :key="index"></v-divider>
    </template>
  </v-list>
</template>
<script>
export default {};
</script>