<script lang="ts">
  import { onMount } from "svelte";
  import { PUBLIC_BACKEND_ADDR, PUBLIC_BACKEND_PORT } from "$env/static/public";
  import { makeBlobURL } from "$lib/imgblob";
  import def_icon from "$lib/assets/def_user_icon.svg";

  type Student = {
    studentID: string
    name: string
    bio: string
    img: string // BLOB URL
  }

  let students: Student[] = $state([])

  onMount(() => {
    fetch(`${PUBLIC_BACKEND_ADDR}:${PUBLIC_BACKEND_PORT}/api/students/profile`)
      .then((res) => res.json())
      .then((data) => {
        data.forEach((student: any) => {
          console.log(student)
          const newStudent: Student = {
            studentID: student.studentID,
            name: student.name,
            bio: student.bio,
            img: def_icon
          }
          if (student.img !== undefined) {
            newStudent.img = makeBlobURL(student.img.content)
          }
          students.push(newStudent)
        })
      });
      console.log(students)
  })
</script>

<div class="margin-case">
  <h1>学生一覧</h1>
  <div id="student-list">
    {#each students as student}
      <div class="student-profile">
        <img src={student.img} alt={student.name} />
        <div class="student-info">
          <h2>{student.name}</h2>
          <span>@{student.studentID}</span>
          <p>{student.bio}</p>
        </div>
      </div>
    {/each}
  </div>
</div>

<style>
  #student-list {
    display: flex;
    flex-flow: row;
    flex-wrap: wrap;
  }
  .student-profile {
    border-radius: 12px;
    background-color: var(--bg-bg-color);
    width: 300px;
    margin: 24px;
    display: flex;
    img {
      margin: 8px;
      width: 100px;
      height: 100px;
    }
    .student-info {
      padding-left: 8px;
      display: flex;
      flex-flow: column;
      span {
        color: var(--immoral-light);
      }
    }
  }
</style>