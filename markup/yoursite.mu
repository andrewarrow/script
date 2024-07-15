div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full flex justify-center
      div text-center text-2xl
        div mt-9 text-white font-mono
          a href=/ link
            script.fly.dev
        div flex flex-wrap mt-3 justify-center
          <svg class="" width="200" height="200" viewBox="0 0 200 200"> <defs> <path id="circlePath" d="M 100, 100 m -50, 0 a 50,50 0 1,1 100,0 a 50,50 0 1,1 -100,0" /> </defs> <text font-size="20" fill="gray"> <textPath href="#circlePath"> yoursite-here.fly.dev </textPath> </text> </svg>
        div mt-0
          Get your site on this site!
        div mt-9
          It's like product hunt, just add a little javascript badge: 
        div flex justify-center
          div
            div mt-6 
              a block href=/ bg-gray-100 rounded-lg w-64 text-sm border border-blue-600
                div p-2 flex space-x-3 items-center
                  div
                    img src=logo.png w-12
                  div space-y-1 text-center w-full
                    div text-gray-900
                      listed on script.fly.dev 
                    div text-gray-900 font-mono
                      rank #5
            div mt-6
              a block href=/ bg-grey-900 rounded-lg w-64 text-sm border border-red-600
                div p-2 flex space-x-3 items-center
                  div
                    img src=logo.png w-12
                  div space-y-1 text-center w-full
                    div 
                      listed on script.fly.dev 
                    div text-gray-100 font-mono 
                      rank #5
        div mt-9
          Use this javascript for dark themes:
        div flex justify-center
          div w-full md:w-1/2 mt-9 font-mono text-sm
            {{.light_script}}
        div mt-9 data-theme=light
          Use this javascript for light themes:
        div flex justify-center data-theme=light
          div w-full md:w-1/2 mt-9 font-mono text-sm
            {{.dark_script}}
