div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div md:w-1/2 flex justify-center
      div
        div 
          form id=stripe text-center text-2xl space-y-3 mt-3
            div
              dev_publishable_key
            div
              input type=text input input-primary id=dev_publishable_key
            div
              dev_secret_key
            div
              input type=text input input-primary id=dev_secret_key
            div
              input id=save type=submit value=Save btn btn-primary btn-sm mt-3
        div 
          {{ range $i, $item := .items }}
            div font-mono
              a href=/core/stripe/{{$item.guid}}
                {{ add $i 1 }}.  {{ $item.dev_publishable_key }}
          {{ end }}
