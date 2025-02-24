<script lang="ts">
  import { goto } from "$app/navigation";
  import { PUBLIC_BACKEND_ADDR, PUBLIC_BACKEND_PORT } from "$env/static/public";
  import { notify } from "$lib/notificationStore";
  import { marked } from "marked";

  let images = $state<Record<string,File>>({})
  let title = $state('')
  let editorStore = $state('# タイトル\n\n本文')

  function removeDefault(e: Event) {
    e.preventDefault()
  }

  function imgUpload(e: DragEvent) {
    e.preventDefault()
    e.stopPropagation()
    
    const files = e.dataTransfer?.files;
    if (!files || files.length === 0) return;
  
    for (const file of files) {
      if (!(file.type.startsWith('image/png') || file.type.startsWith('image/jpeg'))) {
        notify('png, jpegのみアップロード可能です', 'warning')
        continue
      }
      images[file.name] = file
    }
  }
  const renderer = new marked.Renderer()
  renderer.image = ({ href, title, text }) => {
    const dataURL = URL.createObjectURL(images[href])
    return `<img src="${dataURL}" alt="${text}" title="${title || ''}"/>`
  }

  async function upload() {
    if (title === '') {
      notify('記事一覧でのタイトルを付けて下さい', 'warning')
      return
    }
    const fd = new FormData()
    fd.append('articleName', title)
    fd.append('content', editorStore)
    Object.values(images).forEach(value => {
      fd.append('image', value)
    })

    try {
      let res = await fetch(PUBLIC_BACKEND_ADDR + ':' + PUBLIC_BACKEND_PORT + '/api/article/create', {
        method: 'POST',
        body: fd,
      })
      if (res.ok) {
        notify('記事の投稿に成功しました', 'info')
        goto('/')
      } else {
        notify('記事の投稿に失敗しました', 'error')
      }
    } catch (e) {
      notify('記事の投稿に失敗しました(多分サーバが死んでる)', 'error')
    }
  }
</script>

<div class="margin-case">
  <div id="editor-manager">
    <div id="img-folder"
      ondrop={imgUpload}
      ondragover={removeDefault}
      role="region" aria-label="Image upload area"
    >
      {#if Object.keys(images).length > 0}
        <div>
          <p>
          {#each Object.keys(images) as img}
            {img + ' '}
          {/each}
          </p>
        </div>
      {:else}
        <h2>画像フォルダ(ドラッグアンドドロップでアップロード)</h2>
      {/if}
    </div>
    <button onclick={upload}>投稿</button>
  </div>
  <input type="text" placeholder="記事一覧でのタイトル" bind:value={title}>
  <div id="editor-preview-container">
    <div id="editor-container">
      <h2>記事のソース(Markdown)</h2>
      <div bind:innerText={editorStore} id="editor" contenteditable="true">
      </div>
    </div>
    <div id="preview-container">
      <h2>プレビュー</h2>
      <div id="preview">
        {@html marked(editorStore, {renderer})}
      </div>
    </div>
  </div>
</div>

<style>
  .margin-case {
    #editor-manager {
      display: flex;
      flex-flow: row;
      #img-folder {
        padding: 10px;
        height: 30px;
        flex: 1;
        border: 3px dotted var(--text-color);
        display: flex;
        flex-flow: column;
      }
      button {
        width: 100px;
        margin-left: auto;
        background-color: var(--immoral-shadow-darker);
        font-size: 18px;
      }
      #img-folder:hover {
        background-color: var(--immoral-shadow-darker);
        transition: background-color 0.3s ease;
      }
    }
  }
  input {
    width: 100%;
    padding: 10px;
    margin-top: 10px;
    font-size: 18px;
  }
  #editor-preview-container {
    display: flex;
    width: 100%;
    min-height: 100px;

    #editor-container {
      width: 50%;
      #editor {
        border: 1px solid var(--text-color);
        height: 100%;
        padding: 10px;
      }
    }
    #preview-container {
      width: 50%;
      #preview {
        border: 1px solid var(--text-color);
        padding: 10px;
        height: 100%;
        :global(*) {
          width: 100%;
        }
      }
    }
  }
</style>
