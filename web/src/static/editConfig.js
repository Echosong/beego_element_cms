const toolbarOptions = [
    // toggled buttons
    ['bold', 'italic', 'underline', 'strike'],
    ['blockquote'],
    // custom button values
    [{'header': 1}, {'header': 2}],
    [{'list': 'ordered'}, {'list': 'bullet'}],
    // superscript/subscript
    [{'script': 'sub'}, {'script': 'super'}],
    // outdent/indent
    [{'indent': '-1'}, {'indent': '+1'}],
    // text direction
    [{'direction': 'rtl'}],
    [{'size': ['10px', '12px', '14px', '16px', '18px', '20px', '22px', '24px', '26px', '32px', '48px']}],
    // custom dropdown
    [{'header': [1, 2, 3, 4, 5, 6, false]}],
    // dropdown with defaults from theme
    [{'color': []}, {'background': []}],
    [{'font': []}],
    [{'align': []}],
    ['image'],
    // remove formatting button
    ['clean']
  ]
  
  export default {
    toolbarOptions
  }