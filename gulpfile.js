var gulp         = require('gulp'),
    del          = require('del'),
    less         = require('gulp-less'),
    cssmin       = require('gulp-minify-css'),
    browserify   = require('browserify'),
    uglify       = require('gulp-uglify'),
    concat       = require('gulp-concat'),
    source       = require('vinyl-source-stream'),
    buffer       = require('vinyl-buffer'),
    reactify     = require('reactify'),
    glob         = require('glob'),
    environments = require('gulp-environments'),
    pack         = require('./package.json');

var development = environments.development;
var production  = environments.production;

gulp.task('clean', function(cb) {
    del(['./static/css/**', './static/js/**'], cb);
});

gulp.task('bootstrap:css', function() {
    return gulp.src([pack.paths.bootstrap_less])
        .pipe(less())
        .pipe(production(cssmin()))
        .pipe(gulp.dest(pack.dest.css));
});

gulp.task('bootstrap:js', function () {
    return gulp
        .src(pack.paths.bootstrap_js)
        .pipe(gulp.dest(pack.dest.js));
});

gulp.task('bootstrap', ['bootstrap:css', 'bootstrap:js']);

gulp.task('less', function() {
    return gulp.src(pack.paths.less)
        .pipe(less())
        .pipe(concat(pack.dest.less))
        .pipe(gulp.dest(pack.dest.css));
});

gulp.task('less:min', function() {
    return gulp.src(pack.paths.less)
        .pipe(less())
        .pipe(concat(pack.dest.less))
        .pipe(cssmin())
        .pipe(gulp.dest(pack.dest.css));
});


gulp.task('scripts', function() {
    glob(pack.paths.scripts, function(err, files) {
        var opts = {
            entries: files,
            debug:   development()
        };
        
        return browserify(opts)
            .transform(reactify)
            .bundle()
            .pipe(development(source(pack.dest.scripts)))
            .pipe(production(source(pack.dest.scripts)))
            .pipe(buffer())
            .pipe(production(uglify()))
            .pipe(gulp.dest(pack.dest.js));
    });
});

gulp.task('default', ['scripts', 'bootstrap', 'less']);
gulp.task('watch', ['default'], function() {
    return gulp.watch([pack.paths.scripts, pack.paths.less], ['default']);
});

