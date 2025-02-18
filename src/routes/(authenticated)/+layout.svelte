<script>
  import { goto } from '$app/navigation';
  import { notify } from '$lib/notificationStore';
  import { LoggedIn, loadCookie } from '$lib/stores/tokenStore';
  import { onMount } from 'svelte';

  let { children } = $props()
  onMount(() => {
    loadCookie()
    if ($LoggedIn === undefined) {
      notify("学内ページにアクセスするためにはログインが必要です", "warning")
      goto("/login?redirect=/intra")
    }
  })
</script>

{@render children()}
