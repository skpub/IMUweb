<script>
  import { goto } from '$app/navigation';
  import { notify } from '$lib/notificationStore';
  import { LoggedIn } from '$lib/stores/session';
  import { onMount } from 'svelte';

  let { children } = $props()
  onMount(() => {
    if ($LoggedIn === false) {
      notify("学内ページにアクセスするためにはログインが必要です", "warning")
      goto("/login?redirect=/intra")
    }
  })
</script>

{@render children()}
