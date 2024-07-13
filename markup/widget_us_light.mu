{{ define "widget_us_light" }}
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
  {{ end }}
