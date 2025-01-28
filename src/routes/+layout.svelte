<script lang="ts">
  import { goto } from '$app/navigation';
  import imu_logo from '$lib/assets/IMU_logo.svg'
  import imu_text_logo from '$lib/assets/IMU_text_logo.svg'
  import { onMount } from 'svelte';

  let { children } = $props()
  // 0: dark
  // 1: light
  let theme = $state(false)

  $effect(() => {
    if (theme) {
      document.documentElement.classList.add('dark')
      document.documentElement.classList.remove('light')
    } else {
      document.documentElement.classList.add('light')
      document.documentElement.classList.remove('dark')
    }
  })


  onMount(() => {
    const theme_observer_DOM = document.getElementById('is-darkmode')

    const observer = new IntersectionObserver(() => {
      theme = window.matchMedia('(prefers-color-scheme: dark)').matches
    })

    observer.observe(theme_observer_DOM!)
  })

  let contents = [
    {
      name: '学長室',
      link: 'president'
    },
    {
      name: 'キャンパス',
      link: 'campus'
    },
    {
      name: '学部・大学院',
      link: 'faculty'
    },
    {
      name: '入試情報',
      link: 'admission'
    }
  ]
</script>

<header>
    <div id="logo" class="clickable" on:click={() => goto('/')}>
      <img id="imu-logo" src={imu_logo} alt="IMU Logo" />
      <div id="imu-text-logo" style='mask-image: url("{imu_text_logo}");'></div>
    </div>
    <div id="contents">
      {#each contents as content} 
        <a href={`/${content.link}`}>{content.name}</a>
      {/each}
    </div>
</header>
<main>
  {@render children()}
</main>
<div id="is-darkmode"></div>

<style>
  :global(body) {
    margin: 0;
  }
  :global(:root) {
    --immoral-shadow: hsl(275, 8%, 53%);
    --immoral-light: hsl(0, 31%, 88%);
    --immoral-light-darker: hsl(0, 31%, 50%);
    --black: #2a2a2a;
    --white: #f5f5f5;
  }
  :global(a) {
    color: var(--immoral-text);
    text-decoration: none;
  }
  :global(h1, h2, h3) {
    margin: 0;
  }
  :global(.margin-case) {
    margin: 50px;
  }
  :global(.clickable) {
    cursor: pointer;
  }
  :global(.margin-left24) {
    margin-left: 24px;
  }
  :root {
    background: var(--bg-color);
    color-scheme: light dark;
  }
  :root.dark {
    --bg-color: var(--black);
    --text-color: var(--white);
    --immoral-text: var(--immoral-light);
    --key-color: var(--immoral-light);
  }
  :root.light {
    --bg-color: var(--white);
    --text-color: var(--black);
    --immoral-text: var(--immoral-light-darker);
    --key-color: var(--immoral-shadow);
  }
  header {
    display: flex;
    flex-direction: row;
    align-items: center;
    background-color: var(--bg-color);
    height: 90px;
    #logo {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 90px;
      width: 350px;
      #imu-logo {
        width: 70px;
      }
      #imu-text-logo {
        width: 100px;
        height: 100%;
        width: 200px;
        background-color: var(--text-color);
        mask-repeat: no-repeat;
        mask-position: center;
      }
    }
    #contents {
      padding-left: 20px;
      display: flex;
      align-items: center;
      height: 100%;
      flex: 1;
      background: var(--immoral-shadow);
      a {
        margin: 10px;
        color: var(--white);
      }
    }
  }

  /* color theme observer */
  @media (prefers-color-scheme: dark) {
    #is-darkmode {
      position: fixed;
      top: 0;
      left: 0;
      height: 0px;
      width: 0px;
    }
  }
  @media (prefers-color-scheme: light) {
    #is-darkmode {
      display: none;
    }
  }
</style>