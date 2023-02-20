

<template>
 <div id="app">

  <div class="grid-row grid-cols-1">
   <AAlert class="text-[1.75rem]">
     Configured Health Checks 
    </AAlert>
  </div>

  <div class="grid-row grid-cols-4">
    
    <div v-for="check in checks">
      <div v-for="item in check">
        <ACard>
          <div class="a-card-body a-card-spacer">  
            
            <ATypography
              :title="item.Name"
              class="text-lg"
            />
           
            <ATypography
              :subtitle="item.Description"
              class="text-[1.1rem]"
            />
            <text class="text-[0.85rem] text-purple ">Target</text><br>
            <p class="text-[0.85rem]">{{ item.Check.target }}</p>
            <text class="text-[0.85rem] text-purple">Checks</text><br>
            <AChip color="success">{{ item.Check.type }}</AChip><br>
            <text class="text-[0.85rem] text-purple">Actions</text><br>
            <AChip color="info">{{ item.Action.type }}</AChip>
            <br>
              <ASwitch
                v-model="item.Enabled"
                @change="onChange(check)"
                color="info"
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
      checks: [] as any,
      primary: ref(true)
    };
  },
 
  methods: {
    async getConfig() {
    
      const config = await axios.get(apiendpoint + "/getConfig")
      this.checks = config.data.Checks
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



