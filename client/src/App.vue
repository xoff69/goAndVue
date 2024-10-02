<script setup lang="ts">
import { Player } from './model/player';
import axios from 'axios';

import { onMounted } from 'vue'
import { ref} from "vue";

const BASE_URL = 'http://localhost:3002';
const players= ref([]) ;


onMounted(() => {
  console.log(`the component is now mounted.`)
   axios.get(BASE_URL + '/players')
        .then(response => {
          players.value = response.data;
        })
        .catch(error => {
          console.error('There was an error fetching the players :', error);
        });

})



</script>

<template>

<div class="container is-fullhd">
      <table class="table is-fullwidth"  v-if=" players!=null&&players.length > 0">
        <thead  class="thead">
          <tr class="tr">
            <th class="th">Name</th>
            <th class="th">id</th>
          </tr>
        </thead>
        <tbody class="tbody">
          <tr v-for="player in players" :key="player.id">
            
            <td>{{ player.name }}</td>
            <td>{{ player.id }}</td>
          </tr>
        </tbody>
      </table>
      <div  v-if="players!=null&&players.length > 0">
     
    </div>
</div>
</template>
