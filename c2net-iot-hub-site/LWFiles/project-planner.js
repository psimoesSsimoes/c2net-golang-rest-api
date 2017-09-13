
function isValidEmailAddress(emailAddress) {
    var pattern = /^([a-z\d!#$%&'*+\-\/=?^_`{|}~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+(\.[a-z\d!#$%&'*+\-\/=?^_`{|}~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+)*|"((([ \t]*\r\n)?[ \t]+)?([\x01-\x08\x0b\x0c\x0e-\x1f\x7f\x21\x23-\x5b\x5d-\x7e\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|\\[\x01-\x09\x0b\x0c\x0d-\x7f\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]))*(([ \t]*\r\n)?[ \t]+)?")@(([a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|[a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF][a-z\d\-._~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]*[a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])\.)+([a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|[a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF][a-z\d\-._~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]*[a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])\.?$/i;
    return pattern.test(emailAddress);
};

function projectPlanner() {
(function($) {
  var required_message = "This field is required.",
      services = [],
      $form = $('#project--planner-form'),
      ajax_url = $form.attr('action');

  $('input[type="range"]').rangeslider({
		polyfill : false,
		onSlide: function(position, value) {
			var val = Number(value).toLocaleString('en');
			$('.range--placeholder .value').html(val);
		},
	});

  $('input[type="range"]').on('change', function(event) {
    $('input[name="budget-skip"]').val("0");
  });

  $('#agree').on('change', function(event) {
    if ( $(this).prop('checked') == true ) {
      $('.send').removeClass('inactive');
    } else {
      $('.send').addClass('inactive');
    }
  });

  $(".input--span").keydown(function(e){
    var code = e.keyCode || e.which;
    var section = $(this).parents('.project--planner-section'),
        $section_el = $(this).parents('.project--planner-section'),
        input_name = $section_el.find('.input--container').data('input-name');

    if ($(".input--span").text().length > 0) {
      $section_el.find('.error-message').remove();
      $section_el.find('.input--container').removeClass('has-error');
    }

    if ($section_el.find('.input--placeholder-radio.active').length > 0) {
      $section_el.find('.error-message').remove();
    }

    // press enter
		if(code == 13 ) {
      var nextSection = $('.project--planner-section.active').next('.project--planner-section');
      e.preventDefault();
      setTimeout(
  			function() {
  			$(nextSection).find('.note').fadeIn();
  			setTimeout(
  				function() {
  					$(nextSection).find('.note').fadeOut();
  				}, 3000);
  			}, 500);
      $section_el.find('.error-message').remove();
      $section_el.find('.input--container').removeClass('has-error');

      if ( $section_el.find('.input--container.required') && $section_el.find('.input--container.required .input--span:empty').length > 0  ) {
        $section_el.append('<span class="error-message" style="display:none;">'+required_message+'</span>');
        $section_el.find('.error-message').fadeIn(300);
        $section_el.find('.input--container').addClass('has-error');
        return false;
      }
      if ( $section_el.find('.input--container.required') && $section_el.find('.input--placeholder-radio').length > 0 && $section_el.find('.input--placeholder-radio.active').length <= 0) {
        $section_el.append('<span class="error-message" style="display:none;">'+required_message+'</span>');
        $section_el.find('.error-message').fadeIn(300);
        $section_el.find('.input--container').addClass('has-error');
        return false;
      }

      if ( input_name == 'email-address') {
        var email_address = $section_el.find('.input--span').text();
        if( ! isValidEmailAddress( email_address ) ) {
          $section_el.append('<span class="error-message" style="display:none;">Please enter a valid email address.</span>');
          $section_el.find('.error-message').fadeIn(300);
          $section_el.find('.input--container').addClass('has-error');
          return false;
        }
      }

      if ( $section_el.find('.input--container.required') && input_name == 'agree' && $('#agree').prop('checked') == false ) {
          $section_el.append('<span class="error-message" style="display:none;"> Please agree to our Privacy Policy.</span>');
          $section_el.find('.error-message').fadeIn(300);
          return false;
      }

      if ( ! $section_el.is('#project--planner-budget') &&  ! $section_el.is('#project--planner-services') ) {
        $('input[name="'+input_name+'"]').val( $section_el.find('.input--span').text() );
      }

      if ( $section_el.is('#project--planner-services') ) {
        $section_el.find('.input--placeholder-radio.active').each(function(index, el) {
          services.push( $(this).text() );
        });
        $('input[name="'+input_name+'"]').val( services );
      }

      if ($(section).next().is('#project--planner-budget') || $(section).next().is('#project--planner-phone')) {
  			$('.skip').removeClass('hide');$('.nav--note').addClass('hide');
  		} else {
  			$('.skip').addClass('hide');$('.nav--note').removeClass('hide');
  		}

			if ($(section).next('.project--planner-section').length) {
				$(section).removeClass('active');
				$(section).next().fadeIn().addClass('active');
				$('#project--planner-nav a').removeClass('disabled');
			} else {
				e.preventDefault();
        $('#project--planner-nav .send').trigger('click');
			};
			if ($(section).next('.project--planner-section').next('.project--planner-section').length) {
				$('#project--planner-nav .next').removeClass('disabled');
			} else {
				$('#project--planner-nav .next').addClass('disabled send inactive');
			}
    }
	});

	$('.input--placeholder-radio').click(function(){
		$(this).toggleClass('active');
	});

  $('#project--planner-services .input--placeholder-radio').click(function(){
		var service = $(this).html().toLowerCase();
    if ($('#project--planner-services .input--placeholder-radio.active').length > 0) {
      $('#project--planner-services').find('.error-message').remove();
    }
		if ($(this).hasClass('active')) {
			$('#review--services').append('<span id="'+service+'">'+service+'</span>');
		} else {
			$('#'+service).remove();
		}

		var total = $('#review--services span').length;
		$('#review--services span').attr('class','').addClass('service--1of'+total);
	});
  $('.next--section').click(function(){
		var section = $('.project--planner-section.active');
		var nextSection = $('.project--planner-section.active').next('.project--planner-section');
		if ($(section).next('.project--planner-section').length) {
			$(section).removeClass('active');
			$(section).next().fadeIn().addClass('active');
			$('#project--planner-nav a').removeClass('disabled');
		} else {
		}
		$('#project--planner-nav').removeClass('hide');
		setTimeout(
			function() {
			$(nextSection).find('.note').fadeIn();
			setTimeout(
				function() {
					$(nextSection).find('.note').fadeOut();
				}, 3000);
			}, 500);
		if ($(section).next('.project--planner-section').next('.project--planner-section').length) {
			$('#project--planner-nav .next').removeClass('disabled');
		} else {
			$('#project--planner-nav .next').addClass('disabled send inactive');
		}
	});

	$.fn.firstWord = function() {
	  var text = this.text().trim().split(" ");
	  var first = text.shift();
	  return ( text.length > 0 ? "<span class='firstWord'>"+ first + "</span>" : first) ;
	};

  $('#project--planner-nav .next, .skip').click(function(e) {
    e.preventDefault();
		var section = $('.project--planner-section.active'),
        $section_el = $('.project--planner-section.active'),
        input_name = $section_el.find('.input--container').data('input-name');

    $section_el.find('.error-message').remove();
    $section_el.find('.input--container').removeClass('has-error');

    if ( $section_el.find('.input--container.required') && $section_el.find('.input--container.required .input--span:empty').length > 0  ) {
      $section_el.append('<span class="error-message" style="display:none;">'+required_message+'</span>');
      $section_el.find('.error-message').fadeIn(300);
      $section_el.find('.input--container').addClass('has-error');
      return false;
    }
    if ( $section_el.find('.input--container.required') && $section_el.find('.input--placeholder-radio').length > 0 && $section_el.find('.input--placeholder-radio.active').length <= 0) {
      $section_el.append('<span class="error-message" style="display:none;">'+required_message+'</span>');
      $section_el.find('.error-message').fadeIn(300);
      return false;
    }

    if ( input_name == 'email-address') {
      var email_address = $section_el.find('.input--span').text();
      if( ! isValidEmailAddress( email_address ) ) {
        $section_el.append('<span class="error-message" style="display:none;">Please enter a valid email address.</span>');
        $section_el.find('.error-message').fadeIn(300);
        $section_el.find('.input--container').addClass('has-error');
        return false;
      }
    }

    if ( $section_el.find('.input--container.required') && input_name == 'agree' && $('#agree').prop('checked') == false ) {
        $section_el.append('<span class="error-message" style="display:none;"> Please agree to our Privacy Policy.</span>');
        $section_el.find('.error-message').fadeIn(300);
        return false;
    }

    if ( ! $section_el.is('#project--planner-budget') &&  ! $section_el.is('#project--planner-services') ) {
      $('input[name="'+input_name+'"]').val( $section_el.find('.input--span').text() );
    }

    if ( $section_el.is('#project--planner-services') ) {
      $section_el.find('.input--placeholder-radio.active').each(function(index, el) {
        services.push( $(this).text() );
      });
      $('input[name="'+input_name+'"]').val( services );
    }

		if ($(section).next().is('#project--planner-budget') || $(section).next().is('#project--planner-phone')) {
			$('.skip').removeClass('hide');$('.nav--note').addClass('hide');
		} else {
			$('.skip').addClass('hide');$('.nav--note').removeClass('hide');
		}
    if ($(this).hasClass('skip')) {
      $('input[name="budget-skip"]').val("1");
    }
		if ($(this).hasClass('send')) {
      submit_form();
			var $text = $('*[data-input-name="name"]').find('.input--span');
			var name = $text.firstWord();
			$('#review--name').html(name);
			$(section).removeClass('active');
		}
		if ($(section).next('.project--planner-section').length) {
			$(section).removeClass('active');
			$(section).next().fadeIn().addClass('active');
			$('#project--planner-nav a').removeClass('disabled');
		} else {
		}
		if ($(section).next('.project--planner-section').next('.project--planner-section').length) {
			$('#project--planner-nav .next').removeClass('disabled');
      $('#project--planner-nav').removeClass('final');
		} else {
			$('#project--planner-nav .next').addClass('disabled send inactive');
      $('#project--planner-nav').addClass('final');
		}
		setTimeout(
			function() {
			$(section).next('.project--planner-section').find('.note').fadeIn();
			setTimeout(
				function() {
					$(section).next('.project--planner-section').find('.note').fadeOut();
				}, 3000);
			}, 500);
	});

	$('#project--planner-nav .prev').click(function(e) {
    e.preventDefault();
		var section = $('.project--planner-section.active');
		if ($(section).prev().is('#project--planner-budget') || $(section).prev().is('#project--planner-phone')) {
			$('.skip').removeClass('hide');$('.nav--note').addClass('hide');
		} else {
			$('.skip').addClass('hide');$('.nav--note').removeClass('hide');
		}
		if ($(section).prev('.project--planner-section').length) {
			$(section).removeClass('active');
			$(section).prev().fadeIn().addClass('active');
			$('#project--planner-nav a').removeClass('disabled');
		} else {
		}
    $('#project--planner-nav').removeClass('final');
		$('#project--planner-nav a').removeClass('send inactive');
		if ($(section).prev('.project--planner-section').prev('.project--planner-section').length) {
			$('#project--planner-nav').removeClass('hide');
			$('#project--planner-nav .prev').removeClass('disabled');
		} else {
			$('#project--planner-nav').addClass('hide');
			$('#project--planner-nav .prev').addClass('disabled');
		}
	});

	$('.input--placeholder').not('#textarea--toggle').click(function(){
		$(this).hide();
		$(this).siblings('.input--span').get(0).focus();
		$(this).siblings('.input--span').fadeTo( "fast" , 1, function() {
		});
	});

  $('#textarea--toggle').click(function(e){
    e.preventDefault();
    var textarea = $(this).parents('.project--planner-section').children('.textarea--container');
    $(textarea).fadeTo( "fast" , 1, function() {
      $(this).children('textarea').focus();
    });
  });

  var gapage = $('#ga-page').val();

  function submit_form() {
    var ajaxData = new FormData( $form.get( 0 ) );

    $('#project--planner-nav').hide();
    $('.project--planner-review').addClass('active');

    jQuery.ajax({
      action : 'footer_project_planner',
      type: "POST",
      url : ajax_url,
      data: ajaxData,
      cache: false,
      contentType: false,
      processData: false,
      dataType:"json",
      success : function( data ) {
        $('.project--planner-review').removeClass('loading');
        if (data.error == true) {
          $form.find( '.error-message' ).remove();
          $('#project--planner-nav').fadeIn(300);
          $('*[data-input-name="' + data.error_field_id + '"]').parents('.project--planner-section').fadeIn(300).addClass('active').append('<span class="error-message" style="display:none;">' + data.error_message + '</span>');
          $('*[data-input-name="' + data.error_field_id + '"]').parents('.project--planner-section').find('.error-message').fadeIn(300);
        } else {
          $('.ty-message').fadeIn(300);
          ga('send', 'pageview', '/contact-us-success-' + gapage);

          if (gapage === 'southampton') {
            var google_conversion_id = 857444535;
            var google_conversion_language = "en";
            var google_conversion_format = "3";
            var google_conversion_color = "ffffff";
            var google_conversion_label = "90jQCIj_1m8Qt6HumAM";
            var google_remarketing_only = false;
            $.getScript('//www.googleadservices.com/pagead/conversion.js');

            var image = new Image(1, 1);
            image.src = "//www.googleadservices.com/pagead/conversion/857444535/?label=90jQCIj_1m8Qt6HumAM&amp;guid=ON&amp;script=0";
          } else {
            var google_conversion_id = 857444535;
            var google_conversion_language = "en";
            var google_conversion_format = "3";
            var google_conversion_color = "ffffff";
            var google_conversion_label = "5GlMCNGhuG8Qt6HumAM";
            var google_remarketing_only = false;
            $.getScript('//www.googleadservices.com/pagead/conversion.js');

            var image = new Image(1, 1);
            image.src = "//www.googleadservices.com/pagead/conversion/857444535/?label=5GlMCNGhuG8Qt6HumAM&amp;guid=ON&amp;script=0";
          }

          /* Conversion Tracking End */
        }
      },
      error: function() {
        alert( 'Error. Please, contact the webmaster!' );
      }
    });
  }
}(jQuery)); /* end of as page load scripts */
}

jQuery(window).load(function($) {
	projectPlanner();
}(jQuery)); /* end of as page load scripts */
