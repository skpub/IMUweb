<script lang="ts">
  import def_icon from '$lib/assets/def_user_icon.svg'
  import { PUBLIC_BACKEND_ADDR, PUBLIC_BACKEND_PORT } from '$env/static/public';
  import { onMount } from 'svelte';
  import { LoggedIn } from '$lib/stores/tokenStore';
      

  $effect(() => {
    if ($LoggedIn === undefined) return
    fetch(`${PUBLIC_BACKEND_ADDR}:${PUBLIC_BACKEND_PORT}/api/student/profile`, {
      method: 'GET',
      headers: {
        'Authorization': $LoggedIn.token
      }
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data)
      })

  })
</script>