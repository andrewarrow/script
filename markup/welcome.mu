div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div text-center text-2xl
        div mt-9 text-white
          script.fly.dev
        div flex flex-wrap 
          <svg class="" width="200" height="200" viewBox="0 0 200 200"> <defs> <path id="circlePath" d="M 100, 100 m -50, 0 a 50,50 0 1,1 100,0 a 50,50 0 1,1 -100,0" /> </defs> <text font-size="20" fill="white"> <textPath href="#circlePath"> homeducky.com </textPath> </text> </svg>
          <svg class="" width="200" height="200" viewBox="0 0 200 200"> <defs> <path id="circlePath" d="M 100, 100 m -50, 0 a 50,50 0 1,1 100,0 a 50,50 0 1,1 -100,0" /> </defs> <text font-size="20" fill="white"> <textPath href="#circlePath"> g-n-r.fly.dev </textPath> </text> </svg>
          div bg-r w-32 h-32 rounded-full flex justify-center items-center
            div bg-r w-16 h-16 rounded-full 
          div bg-r w-32 h-32 rounded-full flex justify-center items-center
            div bg-r w-16 h-16 rounded-full 
          div bg-r w-32 h-32 rounded-full flex justify-center items-center
            div bg-r w-16 h-16 rounded-full 
          div bg-r w-32 h-32 rounded-full flex justify-center items-center
            div bg-r w-16 h-16 rounded-full 
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
