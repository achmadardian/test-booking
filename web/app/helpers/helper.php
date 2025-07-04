<?php
if (!function_exists('api_url')) {
    function api_url($path = '') {
        return rtrim(config('app.api_url'), '/') . '/' . ltrim($path, '/');
    }
}