// js module
function Promt() {
  // Toast Func
  const toast = function (c) {
    const { msg = "Hello", icon = "success", position = "top-end" } = c;

    const Toast = Swal.mixin({
      toast: true,
      title: msg,
      icon: icon,
      position: position,
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener("mouseenter", Swal.stopTimer);
        toast.addEventListener("mouseleave", Swal.resumeTimer);
      },
    });

    Toast.fire({});
  };

  const success = function (c) {
    const { msg = "Success!", title = "", footer = "" } = c;

    Swal.fire({
      icon: "success",
      title,
      text: msg,
      footer,
    });
  };

  const error = function (c) {
    const { msg = "Ooops!", title = "", footer = "" } = c;

    Swal.fire({
      icon: "error",
      title,
      text: msg,
      footer,
    });
  };

  // multiple input
  async function custom(c) {
    const { msg = "", title = "", showConfirmButton = true } = c;
    let myIcon = ""
    if(c.icon) {
      myIcon = c.icon
    }

    const { value: result } = await Swal.fire({
      icon: myIcon,
      title,
      html: msg,
      backdrop: false,
      focusConfirm: false,
      showCancelButton: true,
      showConfirmButton: showConfirmButton,
      willOpen: () => {
        if (c.willOpen !== undefined) {
          c.willOpen();
        }
      },
      didOpen: () => {
        const startDate = document.getElementById('start');
        const endDate = document.getElementById('end');

        if(startDate && endDate) {
          startDate.removeAttribute('disabled');
          endDate.removeAttribute('disabled');
        }
      },
    });

    if (result) {
      if (result.dismiss !== Swal.DismissReason.cancel) {
        if (result.value != "") {
          if (c.callback !== undefined) {
            c.callback(result);
          }
        } else {
          c.callback(false);
        }
      } else {
        c.callback(false);
      }
    }
  } // end func custom

  // return Promt func
  return {
    toast,
    success,
    error,
    custom,
  };
}