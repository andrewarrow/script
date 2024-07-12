div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div md:w-1/2 flex justify-center mt-9
      div
        div text-white 
          {{.item.dev_publishable_key}}
        div font-mono text-sm
          {{.script}}
