<script>
  import { goto } from "$app/navigation";
  import { notify } from "$lib/notificationStore";
  import { LoggedIn } from "$lib/stores/tokenStore";
  import { onMount } from "svelte";

  onMount(() => {
    if ($LoggedIn === undefined) {
      notify("学内ページにアクセスするためにはログインが必要です", "warning")
      goto("/login?redirect=/intra")
    }
  })
</script>

{#if $LoggedIn !== undefined}
<div class="margin-case">
  <h1>学内ページ</h1>
  <h2>記事作成が出来るよ(右上を見ろ)</h2>
</div>
{/if}
