div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      form id=stripe text-center text-2xl
        div
          dev_publishable_key
        div
          input type=text input input-primary id=dev_publishable_key
        div
          dev_secret_key
        div
          input type=text input input-primary id=dev_secret_key
