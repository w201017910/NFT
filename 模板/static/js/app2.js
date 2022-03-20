/* Website: themesdesign | web and user interface design, development
    Author: Themesdesign
*/

! function($) {
    "use strict";

    var themesd = function() {};

    themesd.prototype.initStickyMenu = function() {
            $(window).scroll(function() {
                var scroll = $(window).scrollTop();

                if (scroll >= 50) {
                    $(".sticky").addClass("nav-sticky");
                } else {
                    $(".sticky").removeClass("nav-sticky");
                }
            });
        },

        themesd.prototype.initSmoothLink = function() {
            $('.navbar-nav a, .anchor-btn').on('click', function(event) {
                var $anchor = $(this);
                $('html, body').stop().animate({
                    scrollTop: $($anchor.attr('href')).offset().top - 50
                }, 1500, 'easeInOutExpo');
                event.preventDefault();
            });
        },

        themesd.prototype.initScrollspy = function() {
            $("#navbarCollapse").scrollspy({
                offset: 50
            });
        },

        // themesd.prototype.initTestimonial = function() {
        //     $("#owl-demo").owlCarousel({
        //         autoPlay: 3000,
        //         stopOnHover: true,
        //         navigation: false,
        //         paginationSpeed: 1000,
        //         goToFirstSpeed: 2000,
        //         singleItem: true,
        //         autoHeight: true,
        //         transitionStyle: "fade"
        //     })
        // },

        themesd.prototype.initContact = function() {
            $('#contact-form').submit(function() {

                var action = $(this).attr('action');

                $("#message").slideUp(750, function() {
                    $('#message').hide();

                    $('#submit')
                        .before('<img src="images/ajax-loader.gif" class="contact-loader" />')
                        .attr('disabled', 'disabled');

                    $.post(action, {
                            name: $('#name').val(),
                            email: $('#email').val(),
                            comments: $('#comments').val(),
                        },
                        function(data) {
                            document.getElementById('message').innerHTML = data;
                            $('#message').slideDown('slow');
                            $('#cform img.contact-loader').fadeOut('slow', function() {
                                $(this).remove()
                            });
                            $('#submit').removeAttr('disabled');
                            if (data.match('success') != null) $('#cform').slideUp('slow');
                        }
                    );

                });

                return false;

            });
        },
        // themesd.prototype.initPortfolio = function() {
        //     //PORTFOLIO FILTER 
        //     var $container = $('.projects-wrapper');
        //     var $filter = $('#filter');
        //     // Initialize isotope 
        //     $container.isotope({
        //         filter: '*',
        //         layoutMode: 'masonry',
        //         animationOptions: {
        //             duration: 750,
        //             easing: 'linear'
        //         }
        //     });
        //     // Filter items when filter link is clicked
        //     $filter.find('a').click(function() {
        //         var selector = $(this).attr('data-filter');
        //         $filter.find('a').removeClass('active');
        //         $(this).addClass('active');
        //         $container.isotope({
        //             filter: selector,
        //             animationOptions: {
        //                 animationDuration: 750,
        //                 easing: 'linear',
        //                 queue: false,
        //             }
        //         });
        //         return false;
        //     });
        //     /*END*/

        // },

        themesd.prototype.init = function() {
            this.initStickyMenu();
            this.initSmoothLink();
            this.initScrollspy();
            // this.initTestimonial();
            this.initContact();
            // this.initPortfolio();
        },

        //init
        $.themesd = new themesd, $.themesd.Constructor = themesd
}(window.jQuery),

//initializing
function($) {
    "use strict";
    $.themesd.init();
}(window.jQuery);


/*global jQuery */
jQuery(function ($) {
    'use strict';

    /**
     * Contact Form Application
     */
    var ContactFormApp = {
        $contactForm: $("#ajax-form"),
        $contactFormBtn: $("#send"),
        $contactFormName: $("#name2"),
        $contactFormEmail: $("#email2"),
        $contactFormMessage: $("#message2"),
        $confirmMessage: $("#ajaxsuccess"),
        $errorMessages: $(".error"),
        $errorName: $("#err-name"),
        $errorEmail: $("#err-emailvld"),
        $errorMessage: $("#err-message"),
        $errorForm: $("#err-form"),
        $errorTimeout: $("#err-timedout"),
        $errorState: $("#err-state"),

        //Validate Contact Us Data
        validate: function () {
            var error = false; // we will set this true if the form isn't valid

            var name = this.$contactFormName.val(); // get the value of the input field
            if(name == "" || name == " " || name == "Name") {
                this.$errorName.show(500);
                this.$errorName.delay(4000);
                this.$errorName.animate({
                    height: 'toggle'  
                }, 500, function() {
                    // Animation complete.
                }); 
                error = true; // change the error state to true
            }

            var email_compare = /^([a-z0-9_.-]+)@([da-z.-]+).([a-z.]{2,6})$/; // Syntax to compare against input
            var email = this.$contactFormEmail.val().toLowerCase(); // get the value of the input field

            if (email == "" || email == " " || email == "E-mail") { // check if the field is empty
                this.$errorEmail.show(500);
                this.$errorEmail.delay(4000);
                this.$errorEmail.animate({
                    height: 'toggle'  
                }, 500, function() {
                    // Animation complete.
                });         
                error = true;
            }
            else if (!email_compare.test(email)) { // if it's not empty check the format against our email_compare variable
                this.$errorEmail.show(500);
                this.$errorEmail.delay(4000);
                this.$errorEmail.animate({
                    height: 'toggle'  
                }, 500, function() {
                    // Animation complete.
                });         
                error = true;
            }

            var message = this.$contactFormMessage.val(); // get the value of the input field
            
            if(message == "" || message == " " || message == "Message") {              
                this.$errorMessage.show(500);
                this.$errorMessage.delay(4000);
                this.$errorMessage.animate({
                    height: 'toggle'  
                }, 500, function() {
                    // Animation complete.
                });            
                error = true; // change the error state to true
            }

            if(error == true) {
                this.$errorForm.show(500);
                this.$errorForm.delay(4000);
                this.$errorForm.animate({
                    height: 'toggle'  
                }, 500, function() {
                    // Animation complete.
                }); 
            }

            return error;
        },
        //contact form submit handler
        contactFormSubmit: function (obj) {
            this.$errorMessages.fadeOut('slow'); // reset the error messages (hides them)

            if(this.validate() == false) {

                var data_string = $('#ajax-form').serialize(); // Collect data from form

                var $this = this;
                $.ajax({
                    type: "POST",
                    url: $this.$contactForm.attr('action'),
                    data: data_string,
                    timeout: 6000,
                    cache: false,
                    crossDomain: false,
                    error: function(request,error) {
                        if (error == "timeout") {
                            $this.$errorTimeout.slideDown('slow');
                        }
                        else {
                            $this.$errorState.slideDown('slow');
                            $this.$errorState.html('An error occurred: ' + error + '');
                        }
                    },
                    success: function() {
                        $this.$confirmMessage.show(500);
                        $this.$confirmMessage.delay(4000);
                        $this.$confirmMessage.animate({
                            height: 'toggle'  
                            }, 500, function() {
                        });    
                        
                        $this.$contactFormName.val('');
                        $this.$contactFormEmail.val('');
                        $this.$contactFormMessage.val('');
                    }
                });
            }
            return false;
        },
        bindEvents: function () {
            //binding submit event
            this.$contactFormBtn.on('click', this.contactFormSubmit.bind(this));
        },
        init: function () {
            //initializing the contact form
            console.log('Contact form is initialized');
            this.bindEvents();
            return this;
        }
    };

    

    //Initializing the app
    ContactFormApp.init({});

});