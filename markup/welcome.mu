div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div text-center text-2xl
        div mt-9 text-white
          script.fly.dev
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
