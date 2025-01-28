<script lang="ts">
  import gakucho_face from '$lib/assets/gakucho_face.webp'
</script>

<div class="margin-case" id="gakucho-aisatsu">
  <h1>学長挨拶</h1>
  <hr>
  <div id="gakucho-aisatsu-contents">
    <div id="gakucho-profile">
      <img id="gakucho-face" src={gakucho_face} alt="">
      <div id="gakucho-profile-text">
        <h1>インモラル大学学長</h1>
        <h1>佐藤海音</h1>
        <p>博士(実践ハラスメント科学)(インモラル大学) ※</p>
        <p style="color: gray">※上記の私のプロフィールは名前と顔以外全て虚偽です。</p>
      </div>
    </div>
    <div id="gakucho-aisatsu-text">
      <p>新入生の皆さん、インモラル大学へようこそ(150dB)。大学全入時代突入と言われて久しいですが、最近になって</p>
      <p>「大学は就職予備校に成り下がっているのではないか。」</p>
      <p>と親御さんや産業界から、はたまたアカデミアからは、</p>
      <p>「成果主義に陥って、基礎研究を蔑ろにして、短期的な利益ばかりを追い求めていないか。」</p>
      <p>と、非難が各所から寄せられております。社会の情勢に応じた変化を大学教育に求めるのは正当なことですが、一方で教育機関であるばかりでなく研究機関でもある大学の役割を鑑みれば、こうした状況に手をこまねいてはいられません。</p>
      <p>「利益を齎すのとは逆に、寧ろ社会に莫大な不利益を齎す大学があってもよいのではないか。」</p>
      <p>この新構想を実現するため、議会前365日爆音デモを挙行し、政治家に不当な圧力を掛けてみだりにみだらに建学したのが、他でもない本学なのです。</p>
      <p>インモラル大学は、、騒音学部、実践ハラスメント学部、そして自称マナー講師学部からなるインモラル学領域特化大学です。ちなみに自称マナー講師学部を卒業すると自称マナー講師の称号が得られます(特に独占業務を行えるわけではありません)。</p>
      <p>コンプライアンスを破壊し、マナーを破壊し、モラルを破壊し、しっかり実践技術を身に着けて、社会に羽ばたいて行ってください。4年後を楽しみにしています。</p>
    </div>
  </div>
</div>

<style>
  #gakucho-aisatsu {
    display: flex;
    flex-flow:column;
    align-items: center;
    #gakucho-aisatsu-contents {
      display: flex;
      justify-content: center;
      #gakucho-aisatsu-text {
        margin: 20px;
      }
    }
    hr {
      width: 100%;
      color: var(--immoral-light);
    }
    #gakucho-profile {
      margin: 20px;
      font-family: serif;
      display: flex;
      flex-flow: column;
      width: 300px;
      #gakucho-profile-text {
        margin-top: 10px;
      }
    }
  }
</style>