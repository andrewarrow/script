div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div md:w-1/2 flex justify-center mt-9
      div
        div text-white 
          {{.item.dev_publishable_key}}
        div mt-3 bg-purple-900 text-white rounded-lg p-3 font-mono text-sm
          {{.script}}
