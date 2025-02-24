<script lang="ts">
  import { goto } from '$app/navigation';
  import imu_logo from '$lib/assets/IMU_logo.svg'
  import imu_text_logo from '$lib/assets/IMU_text_logo.svg'
  import hamburger from '$lib/assets/hamburger.svg'
  import { onMount } from 'svelte';
  import Notification from '$lib/Notification.svelte';
  import { LoggedIn, reload } from '$lib/stores/session';
  import { notify } from '$lib/notificationStore';
  import { MetaTags, deepMerge } from 'svelte-meta-tags';
  import { page } from '$app/state';


  let { data, children } = $props()

  let metaTags = $derived(deepMerge(data.baseMetaTags, page.data.pageMetaTags))
  // 0: dark
  // 1: light
  let theme = $state(false)
  let hamburger_active = $state(false)

  $effect(() => {
    if (theme) {
      document.documentElement.classList.add('dark')
      document.documentElement.classList.remove('light')
    } else {
      document.documentElement.classList.add('light')
      document.documentElement.classList.remove('dark')
    }
  })

  function closeHamburger(e: any) {
    if (e.target.id !== 'hamburger') {
      hamburger_active = false
    }
  }

  onMount(async () => {
    try {
      await reload()
    } catch (e) {
      console.log(e)
    }
    const theme_observer_DOM = document.getElementById('is-darkmode')

    const observer = new IntersectionObserver(() => {
      theme = window.matchMedia('(prefers-color-scheme: dark)').matches
    })

    observer.observe(theme_observer_DOM!)
  })

  let studentContents = [
    {
      name: '記事作成',
      link: 'cms'
    },
    {
      name: 'マイクラサーバ',
      link: 'minecraft'
    },
    {
      name: 'プロフ編集',
      link: 'profile'
    }
  ]

  let contents = [
    {
      name: '学長室',
      link: 'president'
    },
    {
      name: 'アクセス',
      link: 'access'
    },
    {
      name: '学部・大学院',
      link: 'faculty'
    },
    {
      name: '入試情報',
      link: 'admission'
    },
    {
      name: '学生一覧',
      link: 'student'
    },
    {
      name: '学内ページ',
      link: 'intra'
    },
  ]
</script>

<MetaTags {...metaTags} />

<svelte:window onclick={closeHamburger} />

<header>
  <div id="logo" class="clickable" onclick={() => {goto('/'); hamburger_active = false}}>
    <img id="imu-logo" src={imu_logo} alt="IMU Logo" />
    <div id="imu-text-logo" style='mask-image: url("{imu_text_logo}");'></div>
  </div>
  <div id="contents">
    <div id="contents_list_pc">
      <div id="pub-contents">
        {#each contents as content} 
          <a href={`/${content.link}`}>{content.name}</a>
        {/each}
      </div>
      {#if $LoggedIn !== false }
      <div id="student-contents">
        {#each studentContents as content} 
          <a href={`/${content.link}`}>{content.name}</a>
        {/each}
          <span onclick={() => {
            LoggedIn.set(false)
            notify('ログアウトしました', 'info')
          }}>ログアウト</span>
      </div>
      {/if}
    </div>
    <div id="contents_list_mobile">
      <div id="hamburger" class="clickable" style='mask-image: url("{hamburger}");' onclick={() => hamburger_active = !hamburger_active}></div>
      <nav class:active={hamburger_active}>
        {#each contents as content} 
          <a href={`/${content.link}`} onclick={() => hamburger_active = false}>{content.name}</a>
        {/each}
      </nav>
    </div>
  </div>
</header>
<main>
  {@render children()}
</main>
<div id="is-darkmode"></div>
<Notification />

<style>
  :global(body) {
    margin: 0;
  }
  :global(:root) {
    --immoral-shadow: hsl(275, 10%, 45%);
    --immoral-shadow-darker: hsl(275, 10%, 30%);
    --immoral-light: hsl(0, 60%, 78%);
    --immoral-light-darker: hsl(0, 31%, 50%);
    --black: #2a2a2a;
    --blackblack: #252525;
    --white: #f5f5f5;
    --whitewhite: #e9e9e9;
    --emphasis-dark: hsl(0, 100%, 78%);
    --emphasis-light: hsl(0, 100%, 30%);

    --info: hsl(200, 100%, 40%);
    --error: hsl(0, 80%, 40%);
    --warning: hsl(30, 90%, 40%);
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
    --bg-bg-color: var(--blackblack);
    --text-color: var(--white);
    --immoral-text: var(--immoral-light);
    --key-color: var(--immoral-light);
    --emphasis: var(--emphasis-dark);
  }
  :root.light {
    --bg-color: var(--white);
    --bg-bg-color: var(--whitewhite);
    --text-color: var(--black);
    --immoral-text: var(--immoral-light-darker);
    --key-color: var(--immoral-shadow);
    --emphasis: var(--emphasis-light);
  }
  main {
    /* display: flex; */
    /* margin-top: 90px; */
  }
  header {
    display: flex;
    position: sticky;
    position: -webkit-sticky;
    will-change: transform; /* for Safari */
    left: 0;
    top: 0;
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
      /* padding-left: 20px; */
      display: flex;
      align-items: center;
      height: 90px;
      flex: 1;

      a {
        margin: 10px;
        color: var(--white);
      }
    }

    #contents_list_mobile {
      display: flex;
      align-items: center;
      width: 100%;
      height: 100%;

      #hamburger {
        margin: 0 auto;
        width: 30px;
        height: 30px;
        background-color: var(--text-color);
        mask-repeat: no-repeat;
        mask-position: center;
      }

      nav.active {
        display: flex;
        opacity: 90%;
        flex-direction: column;
        position: absolute;
        width: 300px;
        backdrop-filter: blur(10px);
        top: 90px;
        right: 0;
        background-color: var(--immoral-shadow);

        a {
          margin: 10px;
          color: var(--white);
        }
      }

      nav {
        display: none;
      }
    }
  }

  @media (max-width: 768px) {
    #logo {
      position: fixed;
      width: 100dvw;
    }

    #contents {
      background: none;
      position: fixed;
      right: 0;
      width: 90px;

      #contents_list_pc {
        display: none;
      }
    }
  }

  @media (min-width: 769px) {
    #contents {
      background: var(--immoral-shadow);
      #contents_list_mobile {
        display: none;
      }
      #contents_list_pc {
        width: 100%;
        height: 100%;
        display: flex;
        flex-flow: column;
        #pub-contents {
          padding-left: 20px;
          flex: 1;
          display: flex;
          align-items: center;
        }
        #student-contents {
          padding: 4px;
          padding-left: 20px;
          background-color: var(--immoral-shadow-darker);
          span {
            color: var(--white);
          }
        }
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
