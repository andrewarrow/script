{{ define "widget_us_dark" }}
  div mt-6
    a block href=https://script.fly.dev bg-grey-900 rounded-lg w-64 text-sm border border-red-600
      div p-2 flex space-x-3 items-center
        div
          img src=logo.png w-12
        div space-y-1 text-center w-full
          div 
            listed on script.fly.dev 
          div text-gray-100 font-mono 
            rank #5
  {{ end }}
