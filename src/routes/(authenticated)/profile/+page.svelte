<script lang="ts">
  import def_icon from '$lib/assets/def_user_icon.svg'
  import { PUBLIC_BACKEND_ADDR, PUBLIC_BACKEND_PORT } from '$env/static/public';
  import { LoggedIn } from '$lib/stores/tokenStore';
  import { notify } from '$lib/notificationStore';

  let bio = $state('')
  let name = $state('')

  let imgFileInput = $state<HTMLInputElement | null>(null)
  let file = $state<File | null>(null)

  $effect(() => {
    if ($LoggedIn === undefined) return
    fetch(`${PUBLIC_BACKEND_ADDR}:${PUBLIC_BACKEND_PORT}/api/student/profile`, {
      method: 'GET',
      headers: {
        'Authorization': $LoggedIn.token
      }
    })
      .then((res) => res.json())
      .then((data: any) => {
        name = data.name
        bio = data.bio
        const binary = atob(data.img.content)
        const buffer = new Uint8Array(binary.length)
        for (let i = 0; i < binary.length; i++) {
          buffer[i] = binary.charCodeAt(i)
        }
        const blob = new Blob([buffer], { type: "image/jpeg" })
        icon = URL.createObjectURL(blob)
      })
  })

  function updateName() {
    fetch(`${PUBLIC_BACKEND_ADDR}:${PUBLIC_BACKEND_PORT}/api/student/name`, {
      method: 'PUT',
      headers: {
        'Authorization': $LoggedIn!.token,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        name: name
      })
    }).then((res) => res.ok ? notify('名前を更新しました', 'info'): notify('名前の更新に失敗しました', 'error'))
  }

  function updateBio() {
    fetch(`${PUBLIC_BACKEND_ADDR}:${PUBLIC_BACKEND_PORT}/api/student/bio`, {
      method: 'PUT',
      headers: {
        'Authorization' : $LoggedIn!.token,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        bio: bio
      })
    }).then((res) => res.ok ? notify('自己紹介文を更新しました', 'info'): notify('自己紹介文の更新に失敗しました', 'error'))
  }

  async function updateIcon() {
    if (!imgFileInput) {
      notify('画像を選択してください', 'warning')
    }

    const fd = new FormData()
    fd.append('img', file as Blob)

    fetch(`${PUBLIC_BACKEND_ADDR}:${PUBLIC_BACKEND_PORT}/api/student/icon`, {
      method: 'PUT',
      headers: {
        'Authorization': $LoggedIn!.token,
      },
      body: fd
    })
    .then(async (res) => {
      const data = await res.json()
      console.log(data)
    })
    // .then((res) => res.ok ? notify('アイコンを更新しました', 'info'): notify('アイコンの更新に失敗しました', 'error'))
  }

  let icon = $state<string>(def_icon);

  function handleImgFileChange(e: Event) {
    const input = e.target as HTMLInputElement
    if (input.files && input.files[0]) {
      file = input.files[0]
      const reader = new FileReader()
      reader.onload = (e) => {
        if (e.target) {
          icon = e.target.result as string
        }
      }
      reader.readAsDataURL(input.files[0])
    }
  }

  function selectImg(e: Event) {
    imgFileInput?.click()
  }
</script>

<div class="margin-case">
  <div id="student-profile">
    <div id="student-profile-img">
      <button id="select-img" onclick={selectImg}>
        <img src={icon} alt="">
        <input type="file" accept="image/*" bind:this={imgFileInput} onchange={handleImgFileChange} hidden />
      </button>
      <button onclick={updateIcon}>更新</button>
    </div>
    <div id="student-profile-info">
      <div>
        <input bind:value={name} type="text" /> 
        <button onclick={updateName}>更新</button>
      </div>
      <div>
        <textarea bind:value={bio}></textarea>
        <button onclick={updateBio}>更新</button>
      </div>
    </div>
  </div>
</div>

<style>
  #student-profile {
    margin: 0 auto;
    display: flex;
    justify-content: center;
    #student-profile-info {
      display: flex;
      flex-direction: column;
      div {
        margin-bottom: 20px;
        display: flex;
        align-items: end;
        input {
          padding: 5px;
          font-size: 24px;
          width: 300px;
          border: none;
          background-color: var(--bg-bg-color);
          border-bottom: 3px solid var(--immoral-shadow);
        }
        textarea {
          padding: 5px;
          margin: 0;
          border: none;
          border-bottom: 3px solid var(--immoral-shadow);
          background-color: var(--bg-bg-color);
          width: 300px;
          height: 200px;
        }
      }
      button {
        margin-left: 10px;
        background-color: #0000;
        border: none;
        border-bottom: 3px solid var(--immoral-shadow);
      }
    }
    #student-profile-img {
      display: flex;
      flex-flow: column;
      margin-right: 12px;
    }
    #select-img {
      height: 120px;
      width: 120px;
      border: 0;
      img {
        width: 100%;
        height: 100%;
      }
    }
  }
</style>
