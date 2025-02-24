<script>
  import { goto } from '$app/navigation';
  import { notify } from '$lib/notificationStore';
  import { LoggedIn } from '$lib/stores/session';

  let { children } = $props()
  $effect(() => {
    if ($LoggedIn === false) {
      notify("学内ページにアクセスするためにはログインが必要です", "warning")
      goto("/login?redirect=/login")
    }
  })
</script>

{@render children()}
