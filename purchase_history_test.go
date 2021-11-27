package gosteam

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseHTML(t *testing.T) {
	body := `<html class=" responsive" lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <meta name="theme-color" content="#171a21">
    <title>vityarassinskiyyxe's account</title>
    <link rel="shortcut icon" href="/favicon.ico" type="image/x-icon">



    <link href="https://store.akamai.steamstatic.com/public/shared/css/motiva_sans.css?v=Rc2hpzg2Ex3T&amp;l=english" rel="stylesheet" type="text/css" >
    <link href="https://store.akamai.steamstatic.com/public/shared/css/shared_global.css?v=g1l5S_mGcu6S&amp;l=english" rel="stylesheet" type="text/css" >
    <link href="https://store.akamai.steamstatic.com/public/shared/css/buttons.css?v=6PFqex5UPprb&amp;l=english" rel="stylesheet" type="text/css" >
    <link href="https://store.akamai.steamstatic.com/public/css/v6/store.css?v=HfL1D10359S0&amp;l=english" rel="stylesheet" type="text/css" >
    <link href="https://store.akamai.steamstatic.com/public/css/styles_linux.css?v=.srZuTsIiIsbq" rel="stylesheet" type="text/css" >
    <link href="https://store.akamai.steamstatic.com/public/css/v6/account.css?v=P2WLI8B6ddJe&amp;l=english" rel="stylesheet" type="text/css" >
    <link href="https://store.akamai.steamstatic.com/public/css/v6/steamaccount.css?v=3E3qCqtVQSK3&amp;l=english" rel="stylesheet" type="text/css" >
    <link href="https://store.akamai.steamstatic.com/public/shared/css/shared_responsive.css?v=qzosl95jy69-&amp;l=english" rel="stylesheet" type="text/css" >
    <script>
        (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
            (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
            m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
        })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

        ga('create', 'UA-33786258-1', 'auto', {
            'sampleRate': 0.4				});
        ga('set', 'dimension1', true );
        ga('set', 'dimension2', 'External' );
        ga('set', 'dimension3', 'account' );
        ga('set', 'dimension4', "account\/history" );
        ga('send', 'pageview' );

    </script>
    <script type="text/javascript" src="https://store.akamai.steamstatic.com/public/shared/javascript/jquery-1.8.3.min.js?v=.TZ2NKhB-nliU" ></script>
    <script type="text/javascript">$J = jQuery.noConflict();</script><script type="text/javascript">VALVE_PUBLIC_PATH = "https:\/\/store.akamai.steamstatic.com\/public\/";</script><script type="text/javascript" src="https://store.akamai.steamstatic.com/public/shared/javascript/tooltip.js?v=.9Z1XDV02xrml" ></script>

    <script type="text/javascript" src="https://store.akamai.steamstatic.com/public/shared/javascript/shared_global.js?v=BL8X79JqjD-V&amp;l=english" ></script>

    <script type="text/javascript" src="https://store.akamai.steamstatic.com/public/javascript/main.js?v=vx5y-GOBjCpX&amp;l=english" ></script>

    <script type="text/javascript" src="https://store.akamai.steamstatic.com/public/javascript/dynamicstore.js?v=a08q0MGPD3io&amp;l=english" ></script>

    <script type="text/javascript">
        var __PrototypePreserve=[];
        __PrototypePreserve[0] = Array.from;
        __PrototypePreserve[1] = Function.prototype.bind;
    </script>
    <script type="text/javascript" src="https://store.akamai.steamstatic.com/public/javascript/prototype-1.7.js?v=.a38iP7Khdmyy" ></script>
    <script type="text/javascript">
        Array.from = __PrototypePreserve[0] || Array.from;
        Function.prototype.bind = __PrototypePreserve[1] || Function.prototype.bind;
    </script>
    <script type="text/javascript" src="https://store.akamai.steamstatic.com/public/javascript/scriptaculous/_combined.js?v=Me1IBxzktiwk&amp;l=english&amp;load=effects,controls,slider" ></script>

    <script type="text/javascript">Object.seal && [ Object, Array, String, Number ].map( function( builtin ) { Object.seal( builtin.prototype ); } );</script>
    <script type="text/javascript">
        document.addEventListener('DOMContentLoaded', function(event) {
            $J.data( document, 'x_readytime', new Date().getTime() );
            $J.data( document, 'x_oldref', GetNavCookie() );
            SetupTooltips( { tooltipCSSClass: 'store_tooltip'} );
        });
    </script><script type="text/javascript" src="https://store.akamai.steamstatic.com/public/javascript/youraccount.js?v=G4bWsUEfjlTd&amp;l=english" ></script>
    <script type="text/javascript" src="https://store.akamai.steamstatic.com/public/shared/javascript/dselect.js?v=t3m1hHdDkYwJ&amp;l=english" ></script>
    <script type="text/javascript" src="https://store.akamai.steamstatic.com/public/shared/javascript/shared_responsive_adapter.js?v=gcLGc1YQkPQi&amp;l=english" ></script>

    <meta name="twitter:card" content="summary_large_image">

    <meta name="twitter:site" content="@steam" />

    <meta property="og:title" content="vityarassinskiyyxe's account">
    <meta property="twitter:title" content="vityarassinskiyyxe's account">
    <meta property="og:type" content="website">
    <meta property="fb:app_id" content="105386699540688">
    <meta property="og:site" content="Steam">


    <link rel="image_src" href="https://store.akamai.steamstatic.com/public/shared/images/responsive/share_steam_logo.png">
    <meta property="og:image" content="https://store.akamai.steamstatic.com/public/shared/images/responsive/share_steam_logo.png">
    <meta name="twitter:image" content="https://store.akamai.steamstatic.com/public/shared/images/responsive/share_steam_logo.png" />
    <meta property="og:image:secure" content="https://store.akamai.steamstatic.com/public/shared/images/responsive/share_steam_logo.png">





</head>
<body class="v6 account responsive_page">


<div class="responsive_page_frame with_header">
    <div class="responsive_page_menu_ctn mainmenu">
        <div class="responsive_page_menu"  id="responsive_page_menu">
            <div class="mainmenu_contents">
                <div class="mainmenu_contents_items">
                    <!-- profile area -->
                    <div class="responsive_menu_user_area">
                        <div class="responsive_menu_user_persona persona offline">
                            <div class="playerAvatar offline">
                                <a href="https://steamcommunity.com/id/plinius_maior/">
                                    <img src="https://cdn.akamai.steamstatic.com/steamcommunity/public/images/avatars/fe/fef49e7fa7e1997310d705b2a6158ff8dc1cdfeb.jpg" srcset="https://cdn.akamai.steamstatic.com/steamcommunity/public/images/avatars/fe/fef49e7fa7e1997310d705b2a6158ff8dc1cdfeb.jpg 1x, https://cdn.akamai.steamstatic.com/steamcommunity/public/images/avatars/fe/fef49e7fa7e1997310d705b2a6158ff8dc1cdfeb_medium.jpg 2x">											</a>
                            </div>
                            <a href="https://steamcommunity.com/id/plinius_maior/" data-miniprofile="341056105">TRADE BOT#1</a>									</div>

                        <div class="responsive_menu_cartwallet_area persona offline">
                            <div class="responsive_menu_user_cart">
                                <a href="https://store.steampowered.com/cart/?snr=1_account_history__global-responsive-menu">
                                    Cart <b>(0)</b>													</a>
                            </div>
                            <div class="responsive_menu_user_wallet">
                                <a href="https://store.steampowered.com/account/?snr=1_account_history__global-responsive-menu">
                                    Wallet <b>(315,67 pуб.)</b>													</a>
                            </div>
                        </div>
                    </div>

                    <div class="menuitem notifications_item">
                        Notifications										<div class="notification_count_total_ctn has_notifications">
                        <span class="notification_envelope"></span>
                        <span class="notification_count">408</span>
                    </div>
                    </div>
                    <div class="notification_submenu" style="display: none;" data-submenuid="notifications">

                        <a data-notification-type="4" class="popup_menu_item notification_ctn header_notification_comments " href="https://steamcommunity.com/id/plinius_maior/commentnotifications/">
                            <span class="notification_icon"></span>New comments <span class="notification_count">0</span>			</a>
                        <div class="header_notification_dropdown_seperator"></div>
                        <a data-notification-type="5" class="popup_menu_item notification_ctn header_notification_items active_inbox_item" href="https://steamcommunity.com/id/plinius_maior/inventory/">
                            <span class="notification_icon"></span>New items <span class="notification_count">100</span>			</a>
                        <div class="header_notification_dropdown_seperator"></div>
                        <a data-notification-type="6" class="popup_menu_item notification_ctn header_notification_invites active_inbox_item" href="https://steamcommunity.com/id/plinius_maior/home/invites/">
                            <span class="notification_icon"></span>New invites <span class="notification_count">199</span>			</a>

                        <div class="header_notification_dropdown_seperator"></div>
                        <a data-notification-type="8" class="popup_menu_item notification_ctn header_notification_gifts " href="https://steamcommunity.com/id/plinius_maior/inventory/#pending_gifts">
                            <span class="notification_icon"></span>New gifts <span class="notification_count">0</span>			</a>
                        <div class="header_notification_dropdown_seperator"></div>
                        <a data-notification-type="9" class="popup_menu_item notification_ctn header_notification_offlinemessages " href="javascript:void(0)" onclick="LaunchWebChat(); HideMenu( 'header_notification_link', 'header_notification_dropdown' ); return false;">
                            <span class="notification_icon"></span>New messages <span class="notification_count">0</span>			</a>
                        <a data-notification-type="1" class="popup_menu_item notification_ctn hide_when_empty header_notification_tradeoffers active_inbox_item" href="https://steamcommunity.com/id/plinius_maior/tradeoffers/">
                            <div class="header_notification_dropdown_seperator"></div>
                            <span class="notification_icon"></span>New trade offers <span class="notification_count">103</span>					</a>
                        <a data-notification-type="2" class="popup_menu_item notification_ctn hide_when_empty header_notification_asyncgame " href="https://steamcommunity.com/id/plinius_maior/gamenotifications">
                            <div class="header_notification_dropdown_seperator"></div>
                            <span class="notification_icon"></span>New turns waiting <span class="notification_count">0</span>					</a>
                        <a data-notification-type="3" class="popup_menu_item notification_ctn hide_when_empty header_notification_moderatormessage active_inbox_item" href="https://steamcommunity.com/id/plinius_maior/moderatormessages">
                            <div class="header_notification_dropdown_seperator"></div>
                            <span class="notification_icon"></span>New community messages <span class="notification_count">6</span>					</a>
                        <a data-notification-type="10" class="popup_menu_item notification_ctn hide_when_empty header_notification_helprequestreply " href="https://help.steampowered.com/en/wizard/HelpRequests">
                            <div class="header_notification_dropdown_seperator"></div>
                            <span class="notification_icon"></span>Steam Support replies <span class="notification_count">0</span>					</a>
                    </div>
                    <a class="menuitem supernav" href="https://store.steampowered.com/?snr=1_account_history__global-responsive-menu" data-tooltip-type="selector" data-tooltip-content=".submenu_store">
                        Store	</a>
                    <div class="submenu_store" style="display: none;" data-submenuid="store">
                        <a class="submenuitem" href="https://store.steampowered.com/?snr=1_account_history__global-responsive-menu">Home</a>
                        <a class="submenuitem" href="https://store.steampowered.com/explore/?snr=1_account_history__global-responsive-menu">Discovery Queue</a>
                        <a class="submenuitem" href="https://steamcommunity.com/my/wishlist/">Wishlist</a>
                        <a class="submenuitem" href="https://store.steampowered.com/points/shop/?snr=1_account_history__global-responsive-menu">Points Shop</a>
                        <a class="submenuitem" href="https://store.steampowered.com/news/?snr=1_account_history__global-responsive-menu">News</a>
                        <a class="submenuitem" href="https://store.steampowered.com/stats/?snr=1_account_history__global-responsive-menu">Stats</a>
                    </div>


                    <a class="menuitem supernav" style="display: block" href="https://steamcommunity.com/" data-tooltip-type="selector" data-tooltip-content=".submenu_community">
                        Community		</a>
                    <div class="submenu_community" style="display: none;" data-submenuid="community">
                        <a class="submenuitem" href="https://steamcommunity.com/">Home</a>
                        <a class="submenuitem" href="https://steamcommunity.com/discussions/">Discussions</a>
                        <a class="submenuitem" href="https://steamcommunity.com/workshop/">Workshop</a>
                        <a class="submenuitem" href="https://steamcommunity.com/market/">Market</a>
                        <a class="submenuitem" href="https://steamcommunity.com/?subsection=broadcasts">Broadcasts</a>
                    </div>


                    <a class="menuitem supernav username persona_name_text_content" href="https://steamcommunity.com/id/plinius_maior/home/" data-tooltip-type="selector" data-tooltip-content=".submenu_username">
                        You &amp; Friends		</a>
                    <div class="submenu_username" style="display: none;" data-submenuid="username">
                        <a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/home/">Activity</a>
                        <a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/">Profile</a>
                        <a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/friends/">Friends</a>
                        <a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/groups/">Groups</a>			<a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/screenshots/">Content</a>			<a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/badges/">Badges</a>			<a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/inventory/">Inventory</a>		</div>


                    <a class="menuitem" href="https://help.steampowered.com/en/">
                        Support	</a>

                    <div class="minor_menu_items">
                        <a class="menuitem" href="https://store.steampowered.com/account/?snr=1_account_history__global-responsive-menu">Account details</a>
                        <a class="menuitem" href="https://store.steampowered.com/account/preferences?snr=1_account_history__global-responsive-menu">Preferences</a>
                        <div class="menuitem change_language_action">
                            Change language								</div>
                        <div class="menuitem" onclick="Logout();">Change User</div>
                        <div class="menuitem" onclick="Responsive_RequestDesktopView();">
                            View desktop website									</div>
                    </div>
                </div>
                <div class="mainmenu_footer_spacer  "></div>
                <div class="mainmenu_footer">
                    <div class="mainmenu_footer_logo"><img src="https://store.akamai.steamstatic.com/public/shared/images/responsive/logo_valve_footer.png"></div>
                    © Valve Corporation. All rights reserved. All trademarks are property of their respective owners in the US and other countries.								<span class="mainmenu_valve_links">
									<a href="https://store.steampowered.com/privacy_agreement/?snr=1_account_history__global-responsive-menu" target="_blank">Privacy Policy</a>
									&nbsp;| &nbsp;<a href="http://www.valvesoftware.com/legal.htm" target="_blank">Legal</a>
									&nbsp;| &nbsp;<a href="https://store.steampowered.com/subscriber_agreement/?snr=1_account_history__global-responsive-menu" target="_blank">Steam Subscriber Agreement</a>
									&nbsp;| &nbsp;<a href="https://store.steampowered.com/steam_refunds/?snr=1_account_history__global-responsive-menu" target="_blank">Refunds</a>
								</span>
                </div>
            </div>
        </div>
    </div>

    <div class="responsive_local_menu_tab">
    </div>

    <div class="responsive_page_menu_ctn localmenu">
        <div class="responsive_page_menu"  id="responsive_page_local_menu">
            <div class="localmenu_content">
            </div>
        </div>
    </div>



    <div class="responsive_header">
        <div class="responsive_header_content">
            <div id="responsive_menu_logo">
                <img src="https://store.akamai.steamstatic.com/public/shared/images/responsive/header_menu_hamburger.png" height="100%">
                <div class="responsive_header_notification_badge_ctn">
                    <div class="responsive_header_notification_badge notification_count_total_ctn has_notifications">
                        <span class="notification_count">408</span>
                    </div>
                </div>
            </div>
            <div class="responsive_header_logo">
                <a href="https://store.steampowered.com/?snr=1_account_history__global-responsive-menu">
                    <img src="https://store.akamai.steamstatic.com/public/shared/images/responsive/header_logo.png" height="36" border="0" alt="STEAM">
                </a>
            </div>
        </div>
    </div>

    <div class="responsive_page_content_overlay">

    </div>

    <div class="responsive_fixonscroll_ctn nonresponsive_hidden ">
    </div>

    <div class="responsive_page_content">

        <div id="global_header" data-panel="{&quot;flow-children&quot;:&quot;row&quot;}">
            <div class="content">
                <div class="logo">
			<span id="logo_holder">
									<a href="https://store.steampowered.com/?snr=1_account_history__global-header">
						<img src="https://store.akamai.steamstatic.com/public/shared/images/header/logo_steam.svg?t=962016" width="176" height="44">
					</a>
							</span>
                </div>

                <div class="supernav_container">
                    <a class="menuitem supernav" href="https://store.steampowered.com/?snr=1_account_history__global-header" data-tooltip-type="selector" data-tooltip-content=".submenu_store">
                        STORE	</a>
                    <div class="submenu_store" style="display: none;" data-submenuid="store">
                        <a class="submenuitem" href="https://store.steampowered.com/?snr=1_account_history__global-header">Home</a>
                        <a class="submenuitem" href="https://store.steampowered.com/explore/?snr=1_account_history__global-header">Discovery Queue</a>
                        <a class="submenuitem" href="https://steamcommunity.com/my/wishlist/">Wishlist</a>
                        <a class="submenuitem" href="https://store.steampowered.com/points/shop/?snr=1_account_history__global-header">Points Shop</a>
                        <a class="submenuitem" href="https://store.steampowered.com/news/?snr=1_account_history__global-header">News</a>
                        <a class="submenuitem" href="https://store.steampowered.com/stats/?snr=1_account_history__global-header">Stats</a>
                        <a class="submenuitem" href="https://store.steampowered.com/about/?snr=1_account_history__global-header">ABOUT</a>
                    </div>


                    <a class="menuitem supernav" style="display: block" href="https://steamcommunity.com/" data-tooltip-type="selector" data-tooltip-content=".submenu_community">
                        COMMUNITY		</a>
                    <div class="submenu_community" style="display: none;" data-submenuid="community">
                        <a class="submenuitem" href="https://steamcommunity.com/">Home</a>
                        <a class="submenuitem" href="https://steamcommunity.com/discussions/">Discussions</a>
                        <a class="submenuitem" href="https://steamcommunity.com/workshop/">Workshop</a>
                        <a class="submenuitem" href="https://steamcommunity.com/market/">Market</a>
                        <a class="submenuitem" href="https://steamcommunity.com/?subsection=broadcasts">Broadcasts</a>
                    </div>


                    <a class="menuitem supernav username persona_name_text_content" href="https://steamcommunity.com/id/plinius_maior/home/" data-tooltip-type="selector" data-tooltip-content=".submenu_username">
                        TRADE BOT#1		</a>
                    <div class="submenu_username" style="display: none;" data-submenuid="username">
                        <a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/home/">Activity</a>
                        <a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/">Profile</a>
                        <a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/friends/">Friends</a>
                        <a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/groups/">Groups</a>			<a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/screenshots/">Content</a>			<a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/badges/">Badges</a>			<a class="submenuitem" href="https://steamcommunity.com/id/plinius_maior/inventory/">Inventory</a>		</div>

                    <a class="menuitem" href="https://steamcommunity.com/chat/">
                        Chat			</a>

                    <a class="menuitem" href="https://help.steampowered.com/en/">
                        SUPPORT	</a>
                </div>
                <script type="text/javascript">
                    jQuery(function($) {
                        $('#global_header .supernav').v_tooltip({'location':'bottom', 'destroyWhenDone': false, 'tooltipClass': 'supernav_content', 'offsetY':-4, 'offsetX': 1, 'horizontalSnap': 4, 'tooltipParent': '#global_header .supernav_container', 'correctForScreenSize': false});
                    });
                </script>

                <div id="global_actions">
                    <div id="global_action_menu">
                        <div class="header_installsteam_btn header_installsteam_btn_gray">

                            <a class="header_installsteam_btn_content" href="https://store.steampowered.com/about/?snr=1_account_history__global-header">
                                Install Steam						</a>
                        </div>


                        <!-- notification inbox area -->
                        <div id="header_notification_area">
                            <script type="text/javascript">$J(EnableNotificationCountPolling);</script>
                            <div id="header_notification_link" class="header_notification_btn global_header_toggle_button notification_count_total_ctn has_notifications" onclick="ShowMenu( this, 'header_notification_dropdown', 'right' );">
                                <span class="notification_count">408</span>
                                <span class="header_notification_envelope">
				<img src="https://store.akamai.steamstatic.com/public/shared/images/responsive/header_menu_notifications.png" width="11" height="8">
			</span>
                            </div>

                            <div class="popup_block_new" id="header_notification_dropdown" style="display: none;">
                                <div class="popup_body popup_menu">
                                    <a data-notification-type="4" class="popup_menu_item notification_ctn header_notification_comments " href="https://steamcommunity.com/id/plinius_maior/commentnotifications/">
                                        <span class="notification_icon"></span><span class="notification_count_string singular" style="display: none;">1 new comment</span><span class="notification_count_string plural" style=""><span class="notification_count">0</span> new comments</span>			</a>
                                    <div class="header_notification_dropdown_seperator"></div>
                                    <a data-notification-type="5" class="popup_menu_item notification_ctn header_notification_items active_inbox_item" href="https://steamcommunity.com/id/plinius_maior/inventory/">
                                        <span class="notification_icon"></span><span class="notification_count_string singular" style="display: none;">1 new item in your inventory</span><span class="notification_count_string plural" style=""><span class="notification_count">100</span> new items in your inventory</span>			</a>
                                    <div class="header_notification_dropdown_seperator"></div>
                                    <a data-notification-type="6" class="popup_menu_item notification_ctn header_notification_invites active_inbox_item" href="https://steamcommunity.com/id/plinius_maior/home/invites/">
                                        <span class="notification_icon"></span><span class="notification_count_string singular" style="display: none;">1 new invite</span><span class="notification_count_string plural" style=""><span class="notification_count">199</span> new invites</span>			</a>

                                    <div class="header_notification_dropdown_seperator"></div>
                                    <a data-notification-type="8" class="popup_menu_item notification_ctn header_notification_gifts " href="https://steamcommunity.com/id/plinius_maior/inventory/#pending_gifts">
                                        <span class="notification_icon"></span><span class="notification_count_string singular" style="display: none;">1 new gift</span><span class="notification_count_string plural" style=""><span class="notification_count">0</span> new gifts</span>			</a>
                                    <div class="header_notification_dropdown_seperator"></div>
                                    <a data-notification-type="9" class="popup_menu_item notification_ctn header_notification_offlinemessages " href="javascript:void(0)" onclick="LaunchWebChat(); HideMenu( 'header_notification_link', 'header_notification_dropdown' ); return false;">
                                        <span class="notification_icon"></span><span class="notification_count_string singular" style="display: none;">1 unread chat message</span><span class="notification_count_string plural" style=""><span class="notification_count">0</span> unread chat messages</span>			</a>
                                    <a data-notification-type="1" class="popup_menu_item notification_ctn hide_when_empty header_notification_tradeoffers active_inbox_item" href="https://steamcommunity.com/id/plinius_maior/tradeoffers/">
                                        <div class="header_notification_dropdown_seperator"></div>
                                        <span class="notification_icon"></span><span class="notification_count_string singular" style="display: none;">1 new trade notification</span><span class="notification_count_string plural" style=""><span class="notification_count">103</span> new trade notifications</span>					</a>
                                    <a data-notification-type="2" class="popup_menu_item notification_ctn hide_when_empty header_notification_asyncgame " href="https://steamcommunity.com/id/plinius_maior/gamenotifications">
                                        <div class="header_notification_dropdown_seperator"></div>
                                        <span class="notification_icon"></span><span class="notification_count_string singular" style="display: none;">1 turn waiting</span><span class="notification_count_string plural" style=""><span class="notification_count">0</span> new turns waiting</span>					</a>
                                    <a data-notification-type="3" class="popup_menu_item notification_ctn hide_when_empty header_notification_moderatormessage active_inbox_item" href="https://steamcommunity.com/id/plinius_maior/moderatormessages">
                                        <div class="header_notification_dropdown_seperator"></div>
                                        <span class="notification_icon"></span><span class="notification_count_string singular" style="display: none;">1 community message</span><span class="notification_count_string plural" style=""><span class="notification_count">6</span> community messages</span>					</a>
                                    <a data-notification-type="10" class="popup_menu_item notification_ctn hide_when_empty header_notification_helprequestreply " href="https://help.steampowered.com/en/wizard/HelpRequests">
                                        <div class="header_notification_dropdown_seperator"></div>
                                        <span class="notification_icon"></span><span class="notification_count_string singular" style="display: none;">1 reply from Steam Support</span><span class="notification_count_string plural" style=""><span class="notification_count">0</span> replies from Steam Support</span>					</a>
                                </div>
                            </div>
                        </div>
                        <span class="pulldown global_action_link persona_name_text_content" id="account_pulldown" onclick="ShowMenu( this, 'account_dropdown', 'right', 'bottom', true );">
						TRADE BOT#1					</span>
                        <div class="popup_block_new" id="account_dropdown" style="display: none;">
                            <div class="popup_body popup_menu">
                                <a class="popup_menu_item" href="https://steamcommunity.com/id/plinius_maior/">View profile</a>
                                <a class="popup_menu_item" href="https://store.steampowered.com/account/?snr=1_account_history__global-header">Account details</a>
                                <a class="popup_menu_item" href="javascript:Logout();">Logout: <span class="persona online">vityarassinskiyyxe</span></a>
                                <a class="popup_menu_item" href="https://store.steampowered.com/account/preferences/?snr=1_account_history__global-header">Preferences</a>

                                <span class="popup_menu_item" id="account_language_pulldown">Change language</span>
                                <div class="popup_block_new" id="language_dropdown" style="display: none;">
                                    <div class="popup_body popup_menu">
                                        <a class="popup_menu_item tight" href="?l=schinese" onclick="ChangeLanguage( 'schinese' ); return false;">简体中文 (Simplified Chinese)</a>
                                        <a class="popup_menu_item tight" href="?l=tchinese" onclick="ChangeLanguage( 'tchinese' ); return false;">繁體中文 (Traditional Chinese)</a>
                                        <a class="popup_menu_item tight" href="?l=japanese" onclick="ChangeLanguage( 'japanese' ); return false;">日本語 (Japanese)</a>
                                        <a class="popup_menu_item tight" href="?l=koreana" onclick="ChangeLanguage( 'koreana' ); return false;">한국어 (Korean)</a>
                                        <a class="popup_menu_item tight" href="?l=thai" onclick="ChangeLanguage( 'thai' ); return false;">ไทย (Thai)</a>
                                        <a class="popup_menu_item tight" href="?l=bulgarian" onclick="ChangeLanguage( 'bulgarian' ); return false;">Български (Bulgarian)</a>
                                        <a class="popup_menu_item tight" href="?l=czech" onclick="ChangeLanguage( 'czech' ); return false;">Čeština (Czech)</a>
                                        <a class="popup_menu_item tight" href="?l=danish" onclick="ChangeLanguage( 'danish' ); return false;">Dansk (Danish)</a>
                                        <a class="popup_menu_item tight" href="?l=german" onclick="ChangeLanguage( 'german' ); return false;">Deutsch (German)</a>
                                        <a class="popup_menu_item tight" href="?l=english" onclick="ChangeLanguage( 'english' ); return false;">English</a>
                                        <a class="popup_menu_item tight" href="?l=spanish" onclick="ChangeLanguage( 'spanish' ); return false;">Español - España (Spanish - Spain)</a>
                                        <a class="popup_menu_item tight" href="?l=latam" onclick="ChangeLanguage( 'latam' ); return false;">Español - Latinoamérica (Spanish - Latin America)</a>
                                        <a class="popup_menu_item tight" href="?l=greek" onclick="ChangeLanguage( 'greek' ); return false;">Ελληνικά (Greek)</a>
                                        <a class="popup_menu_item tight" href="?l=french" onclick="ChangeLanguage( 'french' ); return false;">Français (French)</a>
                                        <a class="popup_menu_item tight" href="?l=italian" onclick="ChangeLanguage( 'italian' ); return false;">Italiano (Italian)</a>
                                        <a class="popup_menu_item tight" href="?l=hungarian" onclick="ChangeLanguage( 'hungarian' ); return false;">Magyar (Hungarian)</a>
                                        <a class="popup_menu_item tight" href="?l=dutch" onclick="ChangeLanguage( 'dutch' ); return false;">Nederlands (Dutch)</a>
                                        <a class="popup_menu_item tight" href="?l=norwegian" onclick="ChangeLanguage( 'norwegian' ); return false;">Norsk (Norwegian)</a>
                                        <a class="popup_menu_item tight" href="?l=polish" onclick="ChangeLanguage( 'polish' ); return false;">Polski (Polish)</a>
                                        <a class="popup_menu_item tight" href="?l=portuguese" onclick="ChangeLanguage( 'portuguese' ); return false;">Português (Portuguese)</a>
                                        <a class="popup_menu_item tight" href="?l=brazilian" onclick="ChangeLanguage( 'brazilian' ); return false;">Português - Brasil (Portuguese - Brazil)</a>
                                        <a class="popup_menu_item tight" href="?l=romanian" onclick="ChangeLanguage( 'romanian' ); return false;">Română (Romanian)</a>
                                        <a class="popup_menu_item tight" href="?l=russian" onclick="ChangeLanguage( 'russian' ); return false;">Русский (Russian)</a>
                                        <a class="popup_menu_item tight" href="?l=finnish" onclick="ChangeLanguage( 'finnish' ); return false;">Suomi (Finnish)</a>
                                        <a class="popup_menu_item tight" href="?l=swedish" onclick="ChangeLanguage( 'swedish' ); return false;">Svenska (Swedish)</a>
                                        <a class="popup_menu_item tight" href="?l=turkish" onclick="ChangeLanguage( 'turkish' ); return false;">Türkçe (Turkish)</a>
                                        <a class="popup_menu_item tight" href="?l=vietnamese" onclick="ChangeLanguage( 'vietnamese' ); return false;">Tiếng Việt (Vietnamese)</a>
                                        <a class="popup_menu_item tight" href="?l=ukrainian" onclick="ChangeLanguage( 'ukrainian' ); return false;">Українська (Ukrainian)</a>
                                        <a class="popup_menu_item tight" href="http://translation.steampowered.com" target="_blank">Help us translate Steam</a>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <script type="text/javascript">
                            RegisterFlyout( 'account_language_pulldown', 'language_dropdown', 'leftsubmenu', 'bottomsubmenu', true );
                        </script>
                        <div id="header_wallet_ctn">
                            <a class="global_action_link" id="header_wallet_balance" href="https://store.steampowered.com/account/store_transactions/">315,67 pуб.</a>
                        </div>
                    </div>
                    <a href="https://steamcommunity.com/id/plinius_maior/" class="user_avatar playerAvatar offline">
                        <img src="https://cdn.akamai.steamstatic.com/steamcommunity/public/images/avatars/fe/fef49e7fa7e1997310d705b2a6158ff8dc1cdfeb.jpg">
                    </a>
                </div>
            </div>
        </div>
        <div id="responsive_store_nav_ctn"></div><div id="responsive_store_nav_overlay" style="display:none;"><div id="responsive_store_nav_overlay_ctn"></div><div id="responsive_store_nav_overlay_bottom"></div></div><div data-cart-banner-spot="1"></div>
        <div class="responsive_page_template_content" data-panel="{&quot;autoFocus&quot;:true}" >


            <div class="page_header_ctn account_management">
                <script type="text/javascript">
                    var g_AccountID = 341056105;
                    var g_sessionID = "76b684ccda83";
                    var g_ServerTime = 1637917323;

                    $J( InitMiniprofileHovers );


                    GStoreItemData.AddNavParams({
                        __page_default: "1_account_history_",
                        storemenu_recommendedtags: "1_account_history__17"		});
                    GDynamicStore.Init( 341056105, false, "", {"primary_language":0,"secondary_languages":0,"platform_windows":0,"platform_mac":0,"platform_linux":0,"hide_adult_content_violence":0,"hide_adult_content_sex":0,"timestamp_updated":1535313603,"hide_store_broadcast":0,"review_score_preference":0,"timestamp_content_descriptor_preferences_updated":0}, 'RU',
                        {"bNoDefaultDescriptors":false} );
                    GStoreItemData.SetCurrencyFormatter( function( nValueInCents, bWholeUnitsOnly ) { var fmt = function( nValueInCents, bWholeUnitsOnly ) {	var format = v_numberformat( nValueInCents / 100, bWholeUnitsOnly ? 0 : 2, ",", "");format = format.replace( ",00", '' ); return format; };var strNegativeSymbol = '';	if ( nValueInCents < 0 ) { strNegativeSymbol = '-'; nValueInCents = -nValueInCents; }return strNegativeSymbol + fmt( nValueInCents, bWholeUnitsOnly ) + " p\u0443\u0431.";} );
                    GStoreItemData.SetCurrencyMinPriceIncrement( 1 );
                </script>
                <div class="page_content">
                    <div class="breadcrumbs">
                        <div class="blockbg">
                            <a data-panel="{&quot;noFocusRing&quot;:true}" href="https://store.steampowered.com/">Home</a>
                            <span class="breadcrumb_separator">&gt;</span>
                            <a data-panel="{&quot;noFocusRing&quot;:true}" href="https://store.steampowered.com/account/">Account</a>
                            <span class="breadcrumb_separator">&gt;</span>
                            <span class="breadcrumb_current_page">Purchase History</span>
                        </div>
                        <div style="clear: left;"></div>
                    </div>
                    <h2 class="pageheader">vityarassinskiyyxe's Purchase History</h2>
                </div>
            </div>
            <div class="page_content_ctn">
                <div id="main_content" class="page_content" >
                    <div class="wallet_history_click_hint">
                        Problem with a transaction? Select it below to get help.		</div>
                    <table class="wallet_history_table">
                        <thead>
                        <tr>
                            <th rowspan="2" class="wht_date">Date</th>
                            <th rowspan="2" class="wht_items">Items</th>
                            <th rowspan="2" class="wht_type">Type</th>
                            <th rowspan="2" class="wht_total">Total</th>
                            <th class="wht_wallet" colspan="2">Wallet</th>
                        </tr>
                        <tr>
                            <th class="wht_wallet_change">Change</th>
                            <th class="wht_wallet_balance">Balance</th>
                        </tr>
                        </thead>
                        <tbody>
                        <tr data-panel="{&quot;noFocusRing&quot;:true,&quot;focusable&quot;:true,&quot;clickOnActivate&quot;:true}" class="wallet_table_row wallet_table_row_amt_change" onclick="location.href='https://help.steampowered.com/en/wizard/HelpWithMyPurchase?transid=2896343084576892574'">
                            <td class="wht_date">24 Dec, 2020</td>
                            <td data-tooltip-text="Click to get help with this purchase" class="wht_items "">
                            Purchased 23 pуб. Wallet Credit																								</td>
                            <td class="wht_type ">
                                <div>Purchase</div>
                                <div class="wth_payment">
                                    Kiosk																					</div>
                            </td>
                            <td class="wht_total ">
                                23 pуб.					</td>

                            <td class="wht_wallet_change wallet_column" data-tooltip-html="&lt;div class='wallet_tip'&gt;&lt;table&gt;&lt;tr&gt;&lt;td&gt;Previous wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;292,67 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;Change:&lt;/td&gt;&lt;td style='text-align: right'&gt;+23 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;New wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;315,67 pуб.&lt;/td&gt;&lt;/table&gt;&lt;/div&gt;">+23 pуб.</td>
                            <td class="wht_wallet_balance wallet_column" data-tooltip-html="&lt;div class='wallet_tip'&gt;&lt;table&gt;&lt;tr&gt;&lt;td&gt;Previous wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;292,67 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;Change:&lt;/td&gt;&lt;td style='text-align: right'&gt;+23 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;New wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;315,67 pуб.&lt;/td&gt;&lt;/table&gt;&lt;/div&gt;"> 315,67 pуб.</td>
                        </tr>
                        <tr data-panel="{&quot;noFocusRing&quot;:true,&quot;focusable&quot;:true,&quot;clickOnActivate&quot;:true}" class="wallet_table_row wallet_table_row_amt_change" onclick="location.href='https://steamcommunity.com/market/#myhistory'">
                            <td class="wht_date">11 Nov, 2017</td>
                            <td data-tooltip-text="Click to view detailed market history" class="wht_items "">
                            Steam Community Market								</td>
                            <td class="wht_type ">
                                <div>Market Transaction</div>
                                <div class="wth_payment">
                                    Wallet			</div>
                            </td>
                            <td class="wht_total ">
                                <div>
                                    1,06 pуб.				</div>
                            </td>

                            <td class="wht_wallet_change wallet_column" data-tooltip-html="&lt;div class='wallet_tip'&gt;&lt;table&gt;&lt;tr&gt;&lt;td&gt;Previous wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;293,73 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;Change:&lt;/td&gt;&lt;td style='text-align: right'&gt;-1,06 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;New wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;292,67 pуб.&lt;/td&gt;&lt;/table&gt;&lt;/div&gt;">-1,06 pуб.</td>
                            <td class="wht_wallet_balance wallet_column" data-tooltip-html="&lt;div class='wallet_tip'&gt;&lt;table&gt;&lt;tr&gt;&lt;td&gt;Previous wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;293,73 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;Change:&lt;/td&gt;&lt;td style='text-align: right'&gt;-1,06 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;New wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;292,67 pуб.&lt;/td&gt;&lt;/table&gt;&lt;/div&gt;"> 292,67 pуб.</td>
                        </tr>
                        <tr data-panel="{&quot;noFocusRing&quot;:true,&quot;focusable&quot;:true,&quot;clickOnActivate&quot;:true}" class="wallet_table_row wallet_table_row_amt_change" onclick="location.href='https://help.steampowered.com/en/wizard/HelpWithMyPurchase?transid=1210753218179283501'">
                            <td class="wht_date">2 Jun, 2017</td>
                            <td data-tooltip-text="Click to get help with this purchase" class="wht_items "">
                            Purchased 5 pуб. Wallet Credit																								</td>
                            <td class="wht_type ">
                                <div>Purchase</div>
                                <div class="wth_payment">
                                    Kiosk																					</div>
                            </td>
                            <td class="wht_total ">
                                5 pуб.					</td>

                            <td class="wht_wallet_change wallet_column" data-tooltip-html="&lt;div class='wallet_tip'&gt;&lt;table&gt;&lt;tr&gt;&lt;td&gt;Previous wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;288,73 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;Change:&lt;/td&gt;&lt;td style='text-align: right'&gt;+5 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;New wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;293,73 pуб.&lt;/td&gt;&lt;/table&gt;&lt;/div&gt;">+5 pуб.</td>
                            <td class="wht_wallet_balance wallet_column" data-tooltip-html="&lt;div class='wallet_tip'&gt;&lt;table&gt;&lt;tr&gt;&lt;td&gt;Previous wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;288,73 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;Change:&lt;/td&gt;&lt;td style='text-align: right'&gt;+5 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;New wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;293,73 pуб.&lt;/td&gt;&lt;/table&gt;&lt;/div&gt;"> 293,73 pуб.</td>
                        </tr>
                        <tr data-panel="{&quot;noFocusRing&quot;:true,&quot;focusable&quot;:true,&quot;clickOnActivate&quot;:true}" class="wallet_table_row wallet_table_row_amt_change" onclick="location.href='https://steamcommunity.com/market/#myhistory'">
                            <td class="wht_date">9 Apr, 2017</td>
                            <td data-tooltip-text="Click to view detailed market history" class="wht_items "">
                            Steam Community Market								</td>
                            <td class="wht_type ">
                                <div>330 Market Transactions</div>
                                <div class="wth_payment">
                                    Wallet			</div>
                            </td>
                            <td class="wht_total ">
                                <div>
                                    461,27 pуб.				</div>
                            </td>

                            <td class="wht_wallet_change wallet_column" data-tooltip-html="&lt;div class='wallet_tip'&gt;&lt;table&gt;&lt;tr&gt;&lt;td&gt;Previous wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;750 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;Change:&lt;/td&gt;&lt;td style='text-align: right'&gt;-461,27 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;New wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;288,73 pуб.&lt;/td&gt;&lt;/table&gt;&lt;/div&gt;">-461,27 pуб.</td>
                            <td class="wht_wallet_balance wallet_column" data-tooltip-html="&lt;div class='wallet_tip'&gt;&lt;table&gt;&lt;tr&gt;&lt;td&gt;Previous wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;750 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;Change:&lt;/td&gt;&lt;td style='text-align: right'&gt;-461,27 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;New wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;288,73 pуб.&lt;/td&gt;&lt;/table&gt;&lt;/div&gt;"> 288,73 pуб.</td>
                        </tr>
                        <tr data-panel="{&quot;noFocusRing&quot;:true,&quot;focusable&quot;:true,&quot;clickOnActivate&quot;:true}" class="wallet_table_row wallet_table_row_amt_change" onclick="location.href='https://help.steampowered.com/en/wizard/HelpWithMyPurchase?transid=645545582318137886'">
                            <td class="wht_date">9 Apr, 2017</td>
                            <td data-tooltip-text="Click to get help with this purchase" class="wht_items "">
                            Purchased 750 pуб. Wallet Credit																								</td>
                            <td class="wht_type ">
                                <div>Purchase</div>
                                <div class="wth_payment">
                                    QIWI Wallet																					</div>
                            </td>
                            <td class="wht_total ">
                                750 pуб.					</td>

                            <td class="wht_wallet_change wallet_column" data-tooltip-html="&lt;div class='wallet_tip'&gt;&lt;table&gt;&lt;tr&gt;&lt;td&gt;Previous wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;0 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;Change:&lt;/td&gt;&lt;td style='text-align: right'&gt;+750 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;New wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;750 pуб.&lt;/td&gt;&lt;/table&gt;&lt;/div&gt;">+750 pуб.</td>
                            <td class="wht_wallet_balance wallet_column" data-tooltip-html="&lt;div class='wallet_tip'&gt;&lt;table&gt;&lt;tr&gt;&lt;td&gt;Previous wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;0 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;Change:&lt;/td&gt;&lt;td style='text-align: right'&gt;+750 pуб.&lt;/td&gt;&lt;tr&gt;&lt;td&gt;New wallet balance:&lt;/td&gt;&lt;td style='text-align: right'&gt;750 pуб.&lt;/td&gt;&lt;/table&gt;&lt;/div&gt;"> 750 pуб.</td>
                        </tr>
                        <tr data-panel="{&quot;noFocusRing&quot;:true,&quot;focusable&quot;:true,&quot;clickOnActivate&quot;:true}" class="wallet_table_row " onclick="location.href='https://help.steampowered.com/en/wizard/HelpWithTransaction?transid=741220440365058903'">
                            <td class="wht_date">10 Jun, 2016</td>
                            <td data-tooltip-text="Click to get help with this purchase" class="wht_items "">
                            <div style="clear: both">
                                CS:GO Prime Status Upgrade											</div>
                            </td>
                            <td class="wht_type ">
                                <div>Purchase</div>
                                <div class="wth_payment">
                                    QIWI Wallet																					</div>
                            </td>
                            <td class="wht_total ">
                                449 pуб.					</td>

                            <td class="wht_wallet_change"></td>
                            <td class="wht_wallet_balance"></td>
                        </tr>
                        <tr id="more_history" style="display: none"></tr>
                        </tbody>
                    </table>

                    <div class="load_more_history_area">
                        <div data-panel="{&quot;focusable&quot;:true,&quot;clickOnActivate&quot;:true}" id="load_more_button" class="btnv6_blue_hoverfade btn_medium" onclick="WalletHistory_LoadMore(); return false;" style="display:none" >
                            Load More Transactions			</div>
                        <div id="wallet_history_loading" style="display: none; ">
                            <img src="https://store.akamai.steamstatic.com/public/images/login/throbber.gif">
                        </div>
                    </div>
                </div>
                <!-- End Center Column -->
                <div style="clear: both;"></div>


            </div>
        </div>
        <!-- End Main Background -->



        <!-- Footer -->
        <div id="footer_spacer" style="" class=""></div>
        <div id="footer"  class="">
            <div class="footer_content">

                <div class="rule"></div>
                <div id="footer_logo_steam"><img src="https://store.akamai.steamstatic.com/public/images/v6/logo_steam_footer.png" alt="Valve Software" border="0" /></div>

                <div id="footer_logo"><a href="http://www.valvesoftware.com" target="_blank" rel="noreferrer"><img src="https://store.akamai.steamstatic.com/public/images/footerLogo_valve_new.png" alt="Valve Software" border="0" /></a></div>
                <div id="footer_text" data-panel="{&quot;flow-children&quot;:&quot;row&quot;}" >
                    <div>&copy; 2021 Valve Corporation.  All rights reserved.  All trademarks are property of their respective owners in the US and other countries.</div>
                    <div>VAT included in all prices where applicable.&nbsp;&nbsp;

                        <a href="https://store.steampowered.com/privacy_agreement/?snr=1_44_44_" target="_blank" rel="noreferrer">Privacy Policy</a>
                        &nbsp; | &nbsp;
                        <a href="https://store.steampowered.com/legal/?snr=1_44_44_" target="_blank" rel="noreferrer">Legal</a>
                        &nbsp; | &nbsp;
                        <a href="https://store.steampowered.com/subscriber_agreement/?snr=1_44_44_" target="_blank" rel="noreferrer">Steam Subscriber Agreement</a>
                        &nbsp; | &nbsp;
                        <a href="https://store.steampowered.com/steam_refunds/?snr=1_44_44_" target="_blank" rel="noreferrer">Refunds</a>
                        &nbsp; | &nbsp;
                        <a href="https://store.steampowered.com/account/cookiepreferences/?snr=1_44_44_" target="_blank" rel="noreferrer">Cookies</a>

                    </div>
                    <div class="responsive_optin_link">
                        <div class="btn_medium btnv6_grey_black" onclick="Responsive_RequestMobileView()">
                            <span>View mobile website</span>
                        </div>
                    </div>

                </div>



                <div style="clear: left;"></div>
                <br>

                <div class="rule"></div>

                <div class="valve_links" data-panel="{&quot;flow-children&quot;:&quot;row&quot;}" >
                    <a href="http://www.valvesoftware.com/about" target="_blank" rel="noreferrer">About Valve</a>
                    &nbsp; | &nbsp;<a href="http://www.valvesoftware.com" target="_blank" rel="noreferrer">Jobs</a>
                    &nbsp; | &nbsp;<a href="http://www.steampowered.com/steamworks/" target="_blank" rel="noreferrer">Steamworks</a>
                    &nbsp; | &nbsp;<a href="https://partner.steamgames.com/steamdirect" target="_blank" rel="noreferrer">Steam Distribution</a>
                    &nbsp; | &nbsp;<a href="https://help.steampowered.com/en/?snr=1_44_44_">Support</a>
                    &nbsp; | &nbsp;<a href="https://store.steampowered.com/digitalgiftcards/?snr=1_44_44_" target="_blank" rel="noreferrer">Gift Cards</a>
                    &nbsp; | &nbsp;<a href="https://steamcommunity.com/linkfilter/?url=http://www.facebook.com/Steam" target="_blank" rel="noopener"><img src="https://store.akamai.steamstatic.com/public/images/ico/ico_facebook.gif"> Steam</a>
                    &nbsp; | &nbsp;<a href="http://twitter.com/steam" target="_blank" rel="noreferrer"><img src="https://store.akamai.steamstatic.com/public/images/ico/ico_twitter.gif"> @steam</a>
                </div>

            </div>
        </div><!-- End Footer -->

        <script type="text/javascript">
            var g_historyCursor = null;

            function WalletHistory_GetWalletColsToHighlight( target )
            {
                var elems = [ $J(target) ];
                var elemNext = $J(target).next();
                var elemPrev = $J(target).prev();
                if ( elemNext && elemNext.is('td') )
                    elems.push( elemNext );
                else
                    elems.push( elemPrev );

                /* this would highlight the previous wallet balance, if enabled
                    var row = $J(target).parent();
                    var next_row = row.nextAll( 'tr.wallet_table_row_amt_change').first();
                    if ( next_row.is('tr') )
                    {
                        // find the last td
                        elems.push( next_row.children().last() );
                    }
                */
                return elems;
            }

            function WalletHistory_BindTooltips()
            {
                $J('.wallet_column').on('v_tooltip_shown', function( e, toolDiv ) {
                    var cols = WalletHistory_GetWalletColsToHighlight(e.currentTarget);
                    for ( var i = 0; i < cols.length; i++ )
                        $J(cols[i]).addClass('wallet_entry_highlight');
                } );

                $J('.wallet_column').on('v_tooltip_hidden', function( e ) {
                    var cols = WalletHistory_GetWalletColsToHighlight(e.currentTarget);
                    for ( var i = 0; i < cols.length; i++ )
                        $J(cols[i]).removeClass('wallet_entry_highlight');
                } );
            }

            $J( function() {
                WalletHistory_BindTooltips();
            } );

            function WalletHistory_LoadMore()
            {
                $J('#load_more_button').hide();
                if ( g_historyCursor == null )
                    return;

                var request_data = {
                    cursor: g_historyCursor,
                    sessionid: g_sessionID
                };


                g_historyCursor = null;

                $J('#wallet_history_loading').show();
                $J.ajax({
                    type: "POST",
                    url: "https://store.steampowered.com/account/AjaxLoadMoreHistory/",
                    data: request_data
                }).done( function( data ) {
                    if ( data.html )
                    {
                        var elem_prev = $J('#more_history').prev();

                        $J('#more_history').before( data.html );

                        var new_elems = elem_prev.nextAll();
                        new_elems.hide();
                        new_elems.fadeIn( 500 );

                        WalletHistory_BindTooltips();
                    }

                    if ( data.cursor )
                    {
                        g_historyCursor = data.cursor;
                        $J('#load_more_button').fadeIn( 50 );
                    }
                    else
                    {
                        $J('#load_more_button').hide();
                    }
                }).always( function() {
                    $J('#wallet_history_loading').hide();
                } );
            }
        </script>

</body>
</html>`

	reader := strings.NewReader(body)
	fmt.Println(parseHTML(reader))
}
