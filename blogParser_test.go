package main

import (
	"fmt"
	"regexp"
	"testing"
)

var home = `<!DOCTYPE html>
<html lang="zh-cmn-Hans" prefix="og: http://ogp.me/ns#" class="han-init">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1" />
    <title>Corkine's BlOG</title>
    <link rel="stylesheet" href="/assets/vendor/primer-css/css/primer.css">
    <link rel="stylesheet" href="/assets/vendor/primer-markdown/dist/user-content.min.css">
    <link rel="stylesheet" href="/assets/vendor/octicons/octicons/octicons.css">
    <link rel="stylesheet" href="/assets/css/components/collection.css">
    <link rel="stylesheet" href="/assets/css/components/repo-card.css">
    <link rel="stylesheet" href="/assets/css/sections/repo-list.css">
    <link rel="stylesheet" href="/assets/css/sections/mini-repo-list.css">
    <link rel="stylesheet" href="/assets/css/components/boxed-group.css">
    <link rel="stylesheet" href="/assets/css/globals/common.css">
    <link rel="stylesheet" href="/assets/vendor/share.js/dist/css/share.min.css">
    <link rel="stylesheet" href="/assets/css/globals/responsive.css">
    <link rel="stylesheet" href="/assets/css/posts/index.css">
    <!-- Latest compiled and minified CSS -->
    
    <link rel="stylesheet" href="/assets/css/pages/index.css">
    

    
    <link rel="canonical" href="http://blog.mazhangjing.com/">
    <link rel="alternate" type="application/atom+xml" title="Corkine's BlOG" href="/feed.xml">
    <link rel="shortcut icon" href="/favicon.ico">
    
    <meta name="keywords" content="Hello CM,CM,MZJ,mzj,Corkine,corkine,ma zhangjing,mazhangjing,mazj,Mazj">
    <meta name="description" content="这里是 Corkine Ma 的个人博客，用来归类整理我在学习生活中遇到的问题和解决方法，请随意浏览。">
    
    
        
    
    <meta property="og:url" content="http://blog.mazhangjing.com/">
    <meta property="og:site_name" content="Corkine's BlOG">
    <meta property="og:type" content="article">
    <meta property="og:locale" content="zh_CN" />
    
    <script src="/assets/vendor/jquery/dist/jquery.min.js"></script>
    <script src="/assets/js/jquery-ui.js"></script>
    <script type="text/javascript">
    function toggleMenu() {
        var nav = document.getElementsByClassName("site-header-nav")[0];
        if (nav.style.display == "inline-flex") {
          nav.style.display = "none";
        } else {
          nav.style.display = "inline-flex";
        }
    }
    </script>
</head>
<body class="home" data-mz="home">
    <header class="site-header">
        <div class="container">
            <h1><a href="/" title="Corkine's BlOG"><span class="octicon octicon-mark-github"></span> Corkine's BlOG</a></h1>
            <button class="collapsed mobile-visible" type="button" onclick="toggleMenu();">
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
            </button>
            <nav class="site-header-nav" role="navigation">
                
                <a href="/" class=" site-header-nav-item" target="" title="首页">首页</a>
                
                <a href="http://www.mazhangjing.com" class=" site-header-nav-item" target="" title="主站">主站</a>
                
                <a href="http://mu.mazhangjing.com" class=" site-header-nav-item" target="" title="笔记">笔记</a>
                
                <a href="/categories/" class=" site-header-nav-item" target="" title="归档">归档</a>
                
                <a href="http://www.mazhangjing.com/about" class=" site-header-nav-item" target="" title="关于">关于</a>
                
            </nav>
        </div>
    </header>
    <!-- / header -->

    <section class="banner">
    <div class="collection-head">
        <div class="container">
            <div class="collection-title">
              <h1 class="collection-header" id="sub-title"><span>知行合一</span></h1>
                <div class="collection-info">
                    <span class="meta-info mobile-hidden">
                        <span class="octicon octicon-location"></span>
                        Wuhan, China
                    </span>
                    <span class="meta-info">
                        <span class="octicon octicon-organization"></span>
                        <a href="http://www.ccnu.edu.cn" target="_blank">Central China Normal University</a>
                    </span>
                     <span class="meta-info">
                        <span class="octicon octicon-mark-github"></span>
                        <a href="https://github.com/corkine" target="_blank">corkine</a>
                    </span>
                </div>
            </div>
        </div>
    </div>
</section>
<!-- /.banner -->
<section class="container content">
    <div class="columns">
        <div class="column two-thirds" >
            <ol class="repo-list">
                
                <li class="repo-list-item">
                    <h3 class="repo-list-name">
                      <a href="/2021/12/08/sicp-compiler/">SICP 106 - 寄存器机器里的计算</a>
                    </h3>
                    <p class="repo-list-description">
                        本章是 SICP 的最后一章，包含了必要的习题解答。首先使用 Scheme 实现了一个寄存器模拟器，然后手动把一些 Scheme 代码翻译成了汇编代码执行，之后用汇编代码写了一个直接控制的解释器，探讨了这个寄存器机器存储分配的方法并使用汇编代码实现了 LISP 经典的分半垃圾回收机制，最后我们使用 Scheme 代码写了一个编译器，使其可以自动将 Scheme 代码翻译为汇编代码。作为最终的结果，Scheme 代码可以编译执行（比如过程定义），也可以在解释器中解释执行（比如过程调用），这种方式充分利用了编译的性能和解释的开发调试灵活性。
                    </p>
                    <p class="repo-list-meta">
                        <span class="meta-info">
                          <span class="octicon octicon-calendar"></span> 2021/12/08
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Coding" title="Coding">Coding</a>
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Programming" title="Programming">Programming</a>
                        </span>
                        
                    </p>
                </li>
                
                <li class="repo-list-item">
                    <h3 class="repo-list-name">
                      <a href="/2021/11/25/sicp-lang-abstraction/">SICP 105 - 元语言抽象</a>
                    </h3>
                    <p class="repo-list-description">
                        本文是对 Structure and Interpretation of Program - SICP 一书第四章“元数据抽象”的总结，包含了（几乎）所有习题的答案。元数据抽象提供给开发者看待和解决问题的新视角，在领域相关问题上可提供相比较通用程序设计语言更加简洁且具有表现力的解决方案。本文将介绍如何实现一个（不包含错误处理的基于 Scheme 基本过程和控制能力）的基本 Lisp 解释器，并且基于这种解释器的变形实现惰性求值、非确定性问题解决，并且在最后介绍了如何从头实现一门具有基本结构、组合能力和抽象能力的逻辑语言（数据查询语言）。
                    </p>
                    <p class="repo-list-meta">
                        <span class="meta-info">
                          <span class="octicon octicon-calendar"></span> 2021/11/25
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Coding" title="Coding">Coding</a>
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Programming" title="Programming">Programming</a>
                        </span>
                        
                    </p>
                </li>
                
                <li class="repo-list-item">
                    <h3 class="repo-list-name">
                      <a href="/2021/11/10/sicp-status-stream/">SICP 104 - 模块化、状态和对象</a>
                    </h3>
                    <p class="repo-list-description">
                        这是 SICP 第三章的内容总结，包含了全部习题的答案。本文阐述了使用赋值和局部状态实现更好的模块化，探讨了这种模型在模拟真实物体上的优势，以及这种为代码引入“时间维度”带来的后果：对于同一和变化、引用透明性和把握代码的时序性导致的心智负担等问题，尤其在并发模型下的严重水土不服。作为替代，我们引入了同样可以表示状态的流模型，通过对流模型的概念和应用的审视来探索其适用范围，并最终得出结论：当非共享状态大于共享时，对象可更好的实现模块化，反之基于函数式（非赋值）和流模型则更方便。
                    </p>
                    <p class="repo-list-meta">
                        <span class="meta-info">
                          <span class="octicon octicon-calendar"></span> 2021/11/10
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Programming" title="Programming">Programming</a>
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Scheme" title="Scheme">Scheme</a>
                        </span>
                        
                    </p>
                </li>
                
                <li class="repo-list-item">
                    <h3 class="repo-list-name">
                      <a href="/2021/10/31/sicp-data-abstraction/">SICP 103 - 构造数据抽象</a>
                    </h3>
                    <p class="repo-list-description">
                        这是 SICP 第二章的内容总结，包含了大部分习题的答案。数据抽象是软件工程的基石之一，本文阐述了数据抽象的方法，如何对层次性数据进行抽象，如何对数据抽象的不同表现形式进行处理（基于标签、数据导向风格和消息传递风格），如何打造数据抽象层次这几个问题。
                    </p>
                    <p class="repo-list-meta">
                        <span class="meta-info">
                          <span class="octicon octicon-calendar"></span> 2021/10/31
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Programming" title="Programming">Programming</a>
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Scheme" title="Scheme">Scheme</a>
                        </span>
                        
                    </p>
                </li>
                
                <li class="repo-list-item">
                    <h3 class="repo-list-name">
                      <a href="/2021/10/13/sicp-procedure-abstraction/">SICP 102 - 构造过程抽象</a>
                    </h3>
                    <p class="repo-list-description">
                        这是 SICP 第一章的内容总结，包含了所有习题的答案。过程抽象是编程最重要的特性之一，本文阐释了为什么要进行过程抽象，如何进行好过程抽象，过程抽象的主要实践，对不同过程抽象的时间和空间认识，如何进行抽象的抽象，如何基于抽象的抽象进行工作，以及在什么时候需要什么层次的抽象这几个问题。
                    </p>
                    <p class="repo-list-meta">
                        <span class="meta-info">
                          <span class="octicon octicon-calendar"></span> 2021/10/13
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Programming" title="Programming">Programming</a>
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Scheme" title="Scheme">Scheme</a>
                        </span>
                        
                    </p>
                </li>
                
                <li class="repo-list-item">
                    <h3 class="repo-list-name">
                      <a href="/2021/10/12/sicp-101/">SICP 101 - 通过代码看世界</a>
                    </h3>
                    <p class="repo-list-description">
                        本文是对 Structure and Interpretation of Program - SICP 一书导论部分的总结，作者提出程序员往往和三个对象相关：人的大脑、程序集合、计算机本身。SICP 使用 Scheme 语言来教授编程，Scheme 的英文含义正是“做计划”，编程的核心在于对现实世界的不断抽象，如果说艺术解释了我们的梦想，计算机则以程序之名在执行它们。
                    </p>
                    <p class="repo-list-meta">
                        <span class="meta-info">
                          <span class="octicon octicon-calendar"></span> 2021/10/12
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Coding" title="Coding">Coding</a>
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Programming" title="Programming">Programming</a>
                        </span>
                        
                    </p>
                </li>
                
                <li class="repo-list-item">
                    <h3 class="repo-list-name">
                      <a href="/2021/07/01/sdn/">SDN 的过去、现在与未来</a>
                    </h3>
                    <p class="repo-list-description">
                        本文介绍了 Software Define Network (SDN) 的理论和发展，考察了 SDN 落地的现状，提出了一些见解。
                    </p>
                    <p class="repo-list-meta">
                        <span class="meta-info">
                          <span class="octicon octicon-calendar"></span> 2021/07/01
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#SDN" title="SDN">SDN</a>
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#OpenStack" title="OpenStack">OpenStack</a>
                        </span>
                        
                    </p>
                </li>
                
                <li class="repo-list-item">
                    <h3 class="repo-list-name">
                      <a href="/2021/06/30/neutron-management/">OpenStack Neutron 网络实战 (LinuxBridge, OVS)</a>
                    </h3>
                    <p class="repo-list-description">
                        本文介绍了 OpenStack Horizon 界面操作网络相关功能以及操作分别基于 LinuxBridge 和 OVS 的底层实现原理，基于 OpenStack Rocky 和 CentOS 7。
                    </p>
                    <p class="repo-list-meta">
                        <span class="meta-info">
                          <span class="octicon octicon-calendar"></span> 2021/06/30
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#OpenStack" title="OpenStack">OpenStack</a>
                        </span>
                        
                        <span class="meta-info">
                          <span class="octicon octicon-file-directory"></span>
                          <a href="/categories/#Neutron" title="Neutron">Neutron</a>
                        </span>
                        
                    </p>
                </li>
                
            </ol>
        </div>
        <div class="column one-third">
            
<h3>搜索</h3>
<div id="site_search">
    <input type="text" id="search_box" placeholder="Search">
    <button class="btn btn-default" id="site_search_do"><span class="octicon octicon-search"></span></button>
</div>

<ul id="search_results"></ul>

<link rel="stylesheet" type="text/css" href="/assets/css/modules/sidebar-search.css">
<script src="/assets/js/lunr.min.js"></script>
<script src="/assets/js/search.js"></script>



            <h3>我的仓库</h3>



<a href="https://github.com/corkine/pyBook" target="_blank" class="card text-center">
    <div class="thumbnail">
        <div class="card-image geopattern" data-pattern-id="pyBook">
            <div class="card-image-cell">
                <h3 class="card-title">
                    pyBook
                </h3>
            </div>
        </div>
        <div class="caption">
            <div class="card-description">
                <p class="card-text"></p>
            </div>
            <div class="card-text">
                <span class="meta-info" title="10 stars">
                    <span class="octicon octicon-star"></span> 10
                </span>
                <span class="meta-info" title="2 forks">
                    <span class="octicon octicon-git-branch"></span> 2
                </span>
                <span class="meta-info" title="Last updated：2021-10-27 06:02:12 UTC">
                    <span class="octicon octicon-clock"></span>
                    <time datetime="2021-10-27 06:02:12 UTC">2021-10-27</time>
                </span>
            </div>
        </div>
    </div>
</a>
</script>


<a href="https://github.com/corkine/siri-shortcut-assistant" target="_blank" class="card text-center">
    <div class="thumbnail">
        <div class="card-image geopattern" data-pattern-id="siri-shortcut-assistant">
            <div class="card-image-cell">
                <h3 class="card-title">
                    siri-shortcut-assistant
                </h3>
            </div>
        </div>
        <div class="caption">
            <div class="card-description">
                <p class="card-text">A program that uses the Java platform and the mysql database. This program allows you to use the siri shortcut to control your computer and execute cmd commands remotely.</p>
            </div>
            <div class="card-text">
                <span class="meta-info" title="3 stars">
                    <span class="octicon octicon-star"></span> 3
                </span>
                <span class="meta-info" title="0 forks">
                    <span class="octicon octicon-git-branch"></span> 0
                </span>
                <span class="meta-info" title="Last updated：2021-05-10 08:48:08 UTC">
                    <span class="octicon octicon-clock"></span>
                    <time datetime="2021-05-10 08:48:08 UTC">2021-05-10</time>
                </span>
            </div>
        </div>
    </div>
</a>
</script>


<a href="https://github.com/corkine/oneNoteCodeFormatterHelper" target="_blank" class="card text-center">
    <div class="thumbnail">
        <div class="card-image geopattern" data-pattern-id="oneNoteCodeFormatterHelper">
            <div class="card-image-cell">
                <h3 class="card-title">
                    oneNoteCodeFormatterHelper
                </h3>
            </div>
        </div>
        <div class="caption">
            <div class="card-description">
                <p class="card-text">程序用于在将代码片段复制到 OneNote 的时候，保持其空格信息和颜色信息，对于暗黑模式，自动进行反色，适用于从 Web 或者 IDE 复制代码到 OneNote。</p>
            </div>
            <div class="card-text">
                <span class="meta-info" title="3 stars">
                    <span class="octicon octicon-star"></span> 3
                </span>
                <span class="meta-info" title="0 forks">
                    <span class="octicon octicon-git-branch"></span> 0
                </span>
                <span class="meta-info" title="Last updated：2021-12-08 02:31:17 UTC">
                    <span class="octicon octicon-clock"></span>
                    <time datetime="2021-12-08 02:31:17 UTC">2021-12-08</time>
                </span>
            </div>
        </div>
    </div>
</a>
</script>


<a href="https://github.com/corkine/cmBed_Spring" target="_blank" class="card text-center">
    <div class="thumbnail">
        <div class="card-image geopattern" data-pattern-id="cmBed_Spring">
            <div class="card-image-cell">
                <h3 class="card-title">
                    cmBed_Spring
                </h3>
            </div>
        </div>
        <div class="caption">
            <div class="card-description">
                <p class="card-text">小🐴图床后端部分 - 基于 Spring Boot + Scala + Java 实现</p>
            </div>
            <div class="card-text">
                <span class="meta-info" title="3 stars">
                    <span class="octicon octicon-star"></span> 3
                </span>
                <span class="meta-info" title="1 forks">
                    <span class="octicon octicon-git-branch"></span> 1
                </span>
                <span class="meta-info" title="Last updated：2020-04-01 05:52:57 UTC">
                    <span class="octicon octicon-clock"></span>
                    <time datetime="2020-04-01 05:52:57 UTC">2020-04-01</time>
                </span>
            </div>
        </div>
    </div>
</a>
</script>


<a href="https://github.com/corkine/psy4jOld" target="_blank" class="card text-center">
    <div class="thumbnail">
        <div class="card-image geopattern" data-pattern-id="psy4jOld">
            <div class="card-image-cell">
                <h3 class="card-title">
                    psy4jOld
                </h3>
            </div>
        </div>
        <div class="caption">
            <div class="card-description">
                <p class="card-text">A cognitive science package based on object-oriented programming ideas. The program is driven by the JavaFx framework and the JVM platform.</p>
            </div>
            <div class="card-text">
                <span class="meta-info" title="2 stars">
                    <span class="octicon octicon-star"></span> 2
                </span>
                <span class="meta-info" title="0 forks">
                    <span class="octicon octicon-git-branch"></span> 0
                </span>
                <span class="meta-info" title="Last updated：2021-02-13 21:54:45 UTC">
                    <span class="octicon octicon-clock"></span>
                    <time datetime="2021-02-13 21:54:45 UTC">2021-02-13</time>
                </span>
            </div>
        </div>
    </div>
</a>
</script>




        </div>
    </div>
    <div class="pagination text-align">
      <div class="btn-group">
        
            <button disabled="disabled" href="javascript:;" class="btn btn-outline">&laquo;</button>
        
        
            <a href="javascript:;" class="active btn btn-outline">1</a>
        
        
          
              <a href="/page2"  class="btn btn-outline">2</a>
          
        
          
              <a href="/page3"  class="btn btn-outline">3</a>
          
        
          
              <a href="/page4"  class="btn btn-outline">4</a>
          
        
          
              <a href="/page5"  class="btn btn-outline">5</a>
          
        
          
              <a href="/page6"  class="btn btn-outline">6</a>
          
        
          
              <a href="/page7"  class="btn btn-outline">7</a>
          
        
          
              <a href="/page8"  class="btn btn-outline">8</a>
          
        
          
              <a href="/page9"  class="btn btn-outline">9</a>
          
        
          
              <a href="/page10"  class="btn btn-outline">10</a>
          
        
          
              <a href="/page11"  class="btn btn-outline">11</a>
          
        
          
              <a href="/page12"  class="btn btn-outline">12</a>
          
        
          
              <a href="/page13"  class="btn btn-outline">13</a>
          
        
          
              <a href="/page14"  class="btn btn-outline">14</a>
          
        
        
            <a href="/page2"  class="btn btn-outline">&raquo;</a>
        
        </div>
    </div>
    <!-- /pagination -->
</section>
<!-- /section.content -->

    <footer class="container">
        <div class="site-footer" role="contentinfo">
            <div class="copyright left mobile-block">
                    © 2018 Marvin Studio 鄂ICP备17002041号
                    <!--span title="Corkine Ma">Corkine Ma</span> -->
                    <a href="javascript:window.scrollTo(0,0)" class="right mobile-visible">TOP</a>
            </div>

            <ul class="site-footer-links right mobile-hidden">
                <li>
                    <a href="javascript:window.scrollTo(0,0)" >TOP</a>
                </li>
            </ul>
        <!--
			<a href="https://github.com/corkine/corkine.github.io" target="_blank" aria-label="view source code">
                <span class="mega-octicon octicon-mark-github" title="GitHub"></span>
            </a>
			-->
            <ul class="site-footer-links mobile-hidden">
                
                <li>
                    <a href="/" title="首页" target="">首页</a>
                </li>
                
                <li>
                    <a href="http://www.mazhangjing.com" title="主站" target="">主站</a>
                </li>
                
                <li>
                    <a href="http://mu.mazhangjing.com" title="笔记" target="">笔记</a>
                </li>
                
                <li>
                    <a href="/categories/" title="归档" target="">归档</a>
                </li>
                
                <li>
                    <a href="http://www.mazhangjing.com/about" title="关于" target="">关于</a>
                </li>
                
                <li><a href="/feed.xml"><span class="octicon octicon-rss" style="color:orange;"></span></a></li>
            </ul>

        </div>
    </footer>
    <!-- / footer -->
    <script src="/assets/vendor/share.js/dist/js/share.min.js"></script>
    <script src="/assets/js/geopattern.js"></script>
    <script src="/assets/js/prism.js"></script>
    <link rel="stylesheet" href="/assets/css/globals/prism.css">
    <script>
      jQuery(document).ready(function($) {
        // geopattern
        $('.geopattern').each(function(){
          $(this).geopattern($(this).data('pattern-id'));
        });
       // hljs.initHighlightingOnLoad();
      });
    </script>
    
</body>
</html>
`

var url = regexp.MustCompile(`(https?://[A-Za-z0-9./]*)`)

var href = regexp.MustCompile(`(?:href|src)="(/.*?)"`)

var resource = regexp.MustCompile(`(?:href|src)="(/.*?\.(?:css|ico|js|xml))"`)

func TestHome(t *testing.T) {
	submatch := href.FindAllStringSubmatch(home, -1)
	fmt.Printf("%v", submatch)
}

func TestFunc(t *testing.T) {
	findPage([]byte(home))
}

func TestRegex(t *testing.T) {
	submatch := url.FindAllStringSubmatch("hellohttps://baidu.com/work\"", -1)
	fmt.Printf("%v", submatch)
}
