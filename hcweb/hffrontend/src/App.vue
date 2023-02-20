

<template>
 <div id="app">

  <div class="grid-row grid-cols-1">
   <AAlert class="text-[1.75rem]">
     Configured Health Checks 
    </AAlert>
  </div>

  <div class="grid-row grid-cols-4">
    
    <div v-for="records in healthchecks">
      <div v-for="item in records">
        <ACard >
          <div class="a-card-body a-card-spacer">  
            
            <ATypography
              :title="item.Name"
              class="text-lg"
            />
           
            <ATypography
              :subtitle="item.Description"
              class="text-[1.1rem]"
            />

            <text class="text-[0.85rem] text-purple">Checks</text><br>
            <div class="flex mb-4 gap-2"> 
              <!-- <text class="text-[0.85rem] text-purple ">Target</text><br> 
              <p class="text-[0.85rem]">{{ check.target }}</p> -->

              <AChip
                 v-for="check in item.Checks"
                :key="check.type"
                color="success"
              >
                {{ check.type }}
             </AChip>
              
            </div>

            <text class="text-[0.85rem] text-purple">Actions</text><br>
            <div class="flex mb-4 gap-2"> 
              <!-- <text class="text-[0.85rem] text-purple ">Target</text><br> 
              <p class="text-[0.85rem]">{{ check.target }}</p> -->

              <AChip
                 v-for="action in item.Actions"
                :key="action.type"
                color="info"
              >
                {{ action.type }}
             </AChip>
              
            </div>

            <br>
              <ASwitch
                v-model="item.Enabled"
                label="Enable/Disable"
                @change="onChange(records)"
                color="info"
                class="flex mb-4 gap-2"
                />
          </div> 
        </ACard> 
      </div>
    </div>
  </div>
</div>
</template>

<script lang="ts">

import axios from 'axios'
import { ref } from 'vue'

var apiendpoint = "/api/v1"

export default {
  name: "App",
  components: {},
  data() {
    return {
      healthchecks: [] as any,
      primary: ref(true)
    };
  },
 
  methods: {
    async getConfig() {
    
      const config = await axios.get(apiendpoint + "/getConfig")
      this.healthchecks = config.data.Healthchecks
    },
    async onChange(item: any) {
      
      const json = JSON.stringify(item)
      console.log(json)

      const config = await axios.post(apiendpoint + "/setConfig",json,{headers: {'Content-Type': 'application/json'}})
      
    }
  },

  mounted(){this.getConfig()}
}
</script>



