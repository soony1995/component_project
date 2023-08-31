/******************************************
 * My Login
 *
 * Bootstrap 4 Login Page
 *
 * @author          Muhamad Nauval Azhar
 * @uri 			https://nauval.in
 * @copyright       Copyright (c) 2018 Muhamad Nauval Azhar
 * @license         My Login is licensed under the MIT license.
 * @github          https://github.com/nauvalazhar/my-login
 * @version         1.2.0
 *
 * Help me to keep this project alive
 * https://www.buymeacoffee.com/mhdnauvalazhar
 * 
 ******************************************/

'use strict';

$(function() {

	// author badge :)
	var author = '<div style="position: fixed;bottom: 0;right: 20px;background-color: #fff;box-shadow: 0 4px 8px rgba(0,0,0,.05);border-radius: 3px 3px 0 0;font-size: 12px;padding: 5px 10px;">By <a href="https://twitter.com/mhdnauvalazhar">@mhdnauvalazhar</a> &nbsp;&bull;&nbsp; <a href="https://www.buymeacoffee.com/mhdnauvalazhar">Buy me a Coffee</a></div>';
	$("body").append(author);

	$("input[type='password'][data-eye]").each(function(i) {
		var $this = $(this),
			id = 'eye-password-' + i,
			el = $('#' + id);

		$this.wrap($("<div/>", {
			style: 'position:relative',
			id: id
		}));

		$this.css({
			paddingRight: 60
		});
		$this.after($("<div/>", {
			html: 'Show',
			class: 'btn btn-primary btn-sm',
			id: 'passeye-toggle-'+i,
		}).css({
				position: 'absolute',
				right: 10,
				top: ($this.outerHeight() / 2) - 12,
				padding: '2px 7px',
				fontSize: 12,
				cursor: 'pointer',
		}));

		$this.after($("<input/>", {
			type: 'hidden',
			id: 'passeye-' + i
		}));

		var invalid_feedback = $this.parent().parent().find('.invalid-feedback');

		if(invalid_feedback.length) {
			$this.after(invalid_feedback.clone());
		}

		$this.on("keyup paste", function() {
			$("#passeye-"+i).val($(this).val());
		});
		$("#passeye-toggle-"+i).on("click", function() {
			if($this.hasClass("show")) {
				$this.attr('type', 'password');
				$this.removeClass("show");
				$(this).removeClass("btn-outline-primary");
			}else{
				$this.attr('type', 'text');
				$this.val($("#passeye-"+i).val());				
				$this.addClass("show");
				$(this).addClass("btn-outline-primary");
			}
		});
	});

	// $(".my-login-validation").submit(function() {
	// 	var form = $(this);
    //     if (form[0].checkValidity() === false) {
    //       event.preventDefault();
    //       event.stopPropagation();
    //     }
	// 	form.addClass('was-validated');
	// });
	
	$(".my-login-validation").submit(function(event) {
		event.preventDefault(); // 폼의 기본 동작(페이지 새로고침)을 중지
	
		var form = $(this);
	
		// 폼 데이터 가져오기
		var formData = new FormData(form[0]);

		// CORS headers
		var headers = {
			'Access-Control-Allow-Origin': '*', // Replace with the allowed origin
			// Add more CORS headers if needed
		};

		// POST 요청 보내기
		$.ajax({
			type: "POST",
			url: "http://localhost:2222/login",
			data: formData,
			processData: false, // 폼 데이터를 처리하지 않도록 설정
			contentType: false, // 컨텐츠 타입 설정하지 않도록 설정
			headers: headers, // Add the CORS headers here
			success: function(response) {
				// 요청이 성공한 경우 여기에서 처리
				console.log("로그인 성공:", response);
	
				// 원하는 동작을 수행하십시오. 예를 들어, 리디렉션 또는 메시지 표시 등
			},
			error: function(error) {
				// 요청이 실패한 경우 여기에서 처리
				console.error("로그인 실패:", error);
	
				// 원하는 동작을 수행하십시오. 예를 들어, 오류 메시지 표시 등
			}
		});
	});	
});
