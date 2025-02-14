<script lang="ts">
  import { notifications, type OneNotification } from '$lib/notificationStore'

  let currentNotifications: OneNotification[] = []

  $: notifications.subscribe(value => {
    currentNotifications = value
  });
</script>

<div class="notification_container">
  {#each currentNotifications as onenot}
    <div class="notification not_{onenot.type}">
      <p>{onenot.msg}</p>
    </div>
  {/each}
</div>

<style>
  .notification_container {
    color: var(--text-color);
    position: absolute;
    min-width: 200px;
    min-height: 100px;
    /* background-color: var(--background); */
    border-radius: 20px;
    top: 95px;
    right: 10px;
    translate: transform(-100%);
    pointer-events: none;
  }
  .notification {
    height: 40px;
    display: flex;
    align-items: center;
    background-color: var(--emphasis-light);
    color: var(--white);
  }
  p {
    padding-left: 10px;
  }
  :global(.notification_container > .not_info) {
    background-color: var(--info);
  }
  :global(.notification_container > .not_warning) {
    background-color: var(--warning);
  }
  :global(.notification_container > .not_error) {
    background-color: var(--error);
  }
</style>
