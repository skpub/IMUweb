<script>
  import { goto } from '$app/navigation';
  import { notify } from '$lib/notificationStore';
  import { LoggedIn } from '$lib/stores/tokenStore';
  import { onMount } from 'svelte';

  let { children } = $props()
  onMount(() => {
    if ($LoggedIn === undefined) {
      notify("学内ページにアクセスするためにはログインが必要です", "warning")
      goto("/login?redirect=/intra")
    }
  })
</script>

{@render children()}