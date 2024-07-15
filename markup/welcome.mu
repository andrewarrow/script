div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full flex justify-center
      div text-center text-2xl
        div mt-9 text-white font-mono
          script.fly.dev
        div flex flex-wrap mt-3 justify-center
          {{ range $i, $item := .items }}
            a href=https://{{$item.domain}}
              <svg class="" width="200" height="200" viewBox="0 0 200 200"> <defs> <path id="circlePath" d="M 100, 100 m -50, 0 a 50,50 0 1,1 100,0 a 50,50 0 1,1 -100,0" /> </defs> <text font-size="20" fill="gold"> <textPath href="#circlePath"> {{$item.domain}} </textPath> </text> </svg>
          {{ end }}
        div mb-96
        div mt-9
          div mt-3 text-center
            div
              Already have account?
            div
              a btn btn-sm btn-primary href=/core/login
                Login
          div mt-3 
            div
              Need an account?
            div
              a btn btn-sm btn-secondary href=/core/register
                Register
